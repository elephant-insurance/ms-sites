package uf

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

//quoteNumberType type encapsulates functions for generating and interpreting Quote Ids
type quoteNumberType struct{}

var QuoteNumber = &quoteNumberType{}

/*
 * What is a Chomp?
 * A byte is 8 bits. A "nybble" is 4 bits. We need something to deal with 5 bits at a time. So... chomp.
 * A 5-bit number gives 32 possibilities, easily converted to an alphanumeric character, avoiding ambiguous 1/l, 0/O.
 */

/*
 * How this works -- the setup
 * First, decide how many characters long you want your identifier to be (Chomp.idCharLength)
 * Each character gives you five bits of data to work with. So 9 chars => 45 bits.
 * Decide how many of your bits to devote to the timestamp (Chomp.TimestampBits).
 * Timestamps are generated in tenth-miliseconds.
 * There are about 30 bits of tenth-miliseconds in a day.
 * There are a little more than 38 bits worth in a year.
 * 42 bits gives 13.94 years.
 * The rest of your bits will be used for random "salt".
 * More salt will reduce the chance of a collision when generating identifiers on different machines.
 * Pick a date for the epoch. This should be earlier than the earliest day you care about.
 */

/*
 * How this works -- the math
 * To generate a policy ID:
 * 1. Compute the tenth miliseconds (ticks / 1,000) since a certain min date (epoch).
 * 2. Mod this with 2 ^ timestampBits to get an int of the desired length.
 * 3. Salt this int with the necessary number of high-order bits to get the full idCharLength * 5 bits.
 * 4. Break the big int into 5-byte Chomps.
 * 5. Starting with the low-order byte, which changes 10000 times per second,
 *     XOR each byte upward. This allows each byte to change in a big way,
 *     obscuring the sequence. Yield the XOR product for each Chomp.
 * 6. Convert the Chomps into characters and concatenate them into a string.
 *
 * To get a datetime from a policy id:
 * 1. Convert the string into an array of Chomps.
 * 2. Reverse the XOR process to deconfusticate the Chomps.
 * 3. Combine the deconfusticated Chomps into a big integer.
 * 4. Cut off the salt bits to leave a timestampBits-sized integer.
 * 5. Add this integer to the epoch to get back a time.
 */

// maxLegalQuoteNumberLength ...
const maxLegalQuoteNumberLength = 15

// minLegalQuoteNumberLength ...
const minLegalQuoteNumberLength = 8

type chomp struct {
	ByteValue byte
}

//Create is an all-in-one function for generating a Quote Id
func (q *quoteNumberType) Create(timeCreated *time.Time) string {
	timeBits := q.generateTimeBits(timeCreated)
	saltyBits := q.saltToFullBitlength(timeBits)
	chompArray := q.toChomps(saltyBits, 0)
	swirlyChomps := q.confusticate(chompArray)

	return q.chompArrayToChompTokens(swirlyChomps)
}

// Parse extracts date and other data stored in a Quote Id
func (q *quoteNumberType) Parse(quoteIdentifier *string) *time.Time {
	if quoteIdentifier == nil || *quoteIdentifier == "" {
		return nil
	}
	cleanID := ""
	for _, thisChar := range *quoteIdentifier {
		uc := strings.ToUpper(string(thisChar))
		if _, found := idNumbers[uc]; found {
			cleanID += uc
		}
	}
	swirlyChomps := q.fromChompTokens(cleanID)
	chompArray := q.deconfusticate(swirlyChomps)
	saltyBits := q.toNumber(chompArray)
	timeBits := q.unsaltToTimestampBits(saltyBits)

	return q.timeFromBits(timeBits)
}

// idCharLength -How long should the identifier be?
// Each char gives us 5 bits of data to work with.
const idCharLength int32 = 9

// timestampBits -How much of the id should be used to store the time?
// Whatever bits remain will be randomly generated.
// This CANNOT be greater than idCharLength * 5
const timestampBits int32 = 42

//timestampUpperBound is the maximum number of tenth-miliseconds that we can work with
var timestampUpperBound = (&quoteNumberType{}).power(2, timestampBits)

// idChars - Turn a 5-bit number into an alphanumeric character. A, I, O, and U are omitted to
//  increase clarity and reduce the chances of generating an offensive string
var idChars = map[int]string{
	0: "1", 1: "9", 2: "B", 3: "8", 4: "C", 5: "7",
	6: "D", 7: "6", 8: "E", 9: "5", 10: "F", 11: "4",
	12: "G", 13: "3", 14: "H", 15: "2", 16: "Q", 17: "S",
	18: "J", 19: "Z", 20: "K", 21: "Y", 22: "L", 23: "X",
	24: "M", 25: "W", 26: "N", 27: "V", 28: "R", 29: "0",
	30: "P", 31: "T",
}

// idNumbers - Turn a char back into a number.
// Redundant mappings for A, I, O, and U are included for backward compatibility
//  and error tolerance.
var idNumbers = map[string]int{
	"A": 0, "9": 1, "B": 2, "8": 3, "C": 4, "7": 5,
	"D": 6, "6": 7, "E": 8, "5": 9, "F": 10, "4": 11,
	"G": 12, "3": 13, "H": 14, "2": 15, "Q": 16, "S": 17,
	"J": 18, "Z": 19, "K": 20, "Y": 21, "L": 22, "X": 23,
	"M": 24, "W": 25, "N": 26, "V": 27, "R": 28, "U": 29,
	"P": 30, "T": 31, "1": 0, "I": 0, "0": 29, "O": 29,
}

// epoch is the earliest date we care about, expressed as a really big integer.
var epoch = time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)

// filterBits is used to mask away the random bits we add
var filterBits = (&quoteNumberType{}).power(int64(2), (timestampBits+1)) - 1

// Chomp Constructs chomp from a ulong
func (q *quoteNumberType) chomp(seed int64) (*chomp, error) {
	if seed > 31 {
		return nil, fmt.Errorf("Error creating chomp: Seed %v too large", seed)
	}
	rtn := chomp{
		ByteValue: byte(seed),
	}
	return &rtn, nil
}

//xor is The all-important XOR operator
func (c *chomp) xor(c2 *chomp) (*chomp, error) {
	if c == nil || c2 == nil {
		return nil, errors.New("Error attempting to XOR two chomps: arguments cannot be nil")
	}
	xv := byte((*c).ByteValue ^ (*c2).ByteValue)
	rtn := chomp{
		ByteValue: xv,
	}
	return &rtn, nil
}

// takeChomp takes a chomp from the low-order end of a ulong
func (q *quoteNumberType) takeChomp(input int64) (*chomp, int64) {
	bv := input & 31
	removed := chomp{ByteValue: byte(bv)}
	remain := int64(0)
	if input > 31 {
		remain = input >> 5
	}
	return &removed, remain
}

// toChomps breaks an int64 into a big-endian array of chomps
// pads the array to a determined length
func (q *quoteNumberType) toChomps(input int64, pad int32) []*chomp {
	if pad == 0 {
		pad = idCharLength
	}
	var chompList []*chomp
	remainder := input
	for remainder > 0 {
		var bittenOff *chomp
		bittenOff, remainder = q.takeChomp(remainder)
		chompList = append(chompList, bittenOff)
	}

	for int32(len(chompList)) < pad {
		mtchomp := chomp{ByteValue: byte(0)}
		chompList = append(chompList, &mtchomp)
	}

	length := len(chompList)
	rtnList := make([]*chomp, length)
	for idx := 0; idx < length; idx++ {
		rtnList[idx] = chompList[length-idx-1]
	}

	return rtnList
}

//toNumber converts a big-endian collection of chomps into an int64
func (q *quoteNumberType) toNumber(chomps []*chomp) int64 {
	rtn := int64(0)
	for _, thisChomp := range chomps {
		rtn = int64((rtn << 5) | int64(thisChomp.ByteValue))
	}
	return rtn
}

//toToken converts a chomp into a token character
func (c *chomp) toToken() string {
	rtn, _ := idChars[int(c.ByteValue)]
	return rtn
}

//fromToken converts a token character into a chomp
func (c *chomp) fromToken(input string) error {
	key := strings.ToUpper(strings.TrimSpace(input))
	if len(key) > 1 {
		return fmt.Errorf("Error creating chomp from token. Must be a single character: %v", input)
	}

	if thisByte, found := idNumbers[key]; !found {
		return fmt.Errorf("Error creating chomp from token. Invalid token: %v", input)
	} else {
		c.ByteValue = byte(thisByte)
	}
	return nil
}

//intToChompTokens turns a big integer into a string of chomp tokens
func (q *quoteNumberType) intToChompTokens(input int64) (string, error) {
	if input < 0 {
		return "", fmt.Errorf("Error creating chomp tokens. Must be nonnegative: %v", input)
	}

	chomps := q.toChomps(input, 0)
	return q.chompArrayToChompTokens(chomps), nil
}

//chompArrayToChompTokens turns an array of chomps into a string
func (q *quoteNumberType) chompArrayToChompTokens(src []*chomp) string {
	rtn := ""
	for _, thisChomp := range src {
		rtn += thisChomp.toToken()
	}
	return rtn
}

//fromChompTokens turns an ID string into an array of chomps
func (q *quoteNumberType) fromChompTokens(input string) []*chomp {
	trimPut := strings.ToUpper(strings.TrimSpace(input))
	rtn := make([]*chomp, len(trimPut))
	for idx, thisChar := range trimPut {
		thisChomp := chomp{}
		thisChomp.fromToken(string(thisChar))
		rtn[idx] = &thisChomp
	}

	return rtn
}

//numberfromChompTokens turns a Quote ID into an integer
func (q *quoteNumberType) numberfromChompTokens(input string) int64 {
	chomps := q.fromChompTokens(input)

	return q.toNumber(chomps)
}

//confusticate adds cyclic xor obfuscation to an array of chomps
func (q *quoteNumberType) confusticate(src []*chomp) []*chomp {
	length := len(src)
	first := src[0]
	last := src[length-1]
	rtn := make([]*chomp, length)
	lastChomp, _ := first.xor(last)
	rtn[0] = lastChomp

	for cnt := 1; cnt < length; cnt++ {
		rtn[cnt], _ = rtn[cnt-1].xor(src[cnt])
	}

	return rtn
}

//deconfusticate removes cyclic xor obfuscation from an array of chomps
func (q *quoteNumberType) deconfusticate(src []*chomp) []*chomp {
	length := len(src)
	rtn := make([]*chomp, length)
	secondToLast := src[length-2]
	last := src[length-1]
	lastChomp, _ := last.xor(secondToLast)
	rtn[0], _ = lastChomp.xor(src[0])

	for idx := 1; idx < length; idx++ {
		rtn[idx], _ = src[idx-1].xor(src[idx])
	}

	return rtn
}

//toString Converts a chomp to a string
func (c *chomp) toString() string {
	return fmt.Sprintf("%v", c.ByteValue)
}

// generateTimeBits creates a big integer from a given time or current time
func (q *quoteNumberType) generateTimeBits(date *time.Time) int64 {
	if date == nil {
		// sleep for a tenth-milisecond to make sure we aren't hitting this multiple times too quickly
		time.Sleep(100000 * time.Nanosecond)
		jetzt := time.Now()
		date = &jetzt
	}

	// returns the number of tenth-seconds since the epoch, modded to bitlength
	dur := date.Sub(epoch)
	return (int64(dur.Nanoseconds() / 100000)) & filterBits
}

// saltToFullBitlength adds random high-order bits to a ulong to yield a longer ulong
func (q *quoteNumberType) saltToFullBitlength(src int64) int64 {
	maxRand := q.power(2, ((idCharLength * 5) - timestampBits))
	salt := rand.Int63n(maxRand)

	return (salt * timestampUpperBound) + src
}

// unsaltToTimestampBits removes random salt bits from a QuoteNumber
func (q *quoteNumberType) unsaltToTimestampBits(src int64) int64 {
	return src % timestampUpperBound
}

// timeFromBits extracts an actual time from a computed integer
func (q *quoteNumberType) timeFromBits(bits int64) *time.Time {
	for true {
		dur := time.Duration(time.Nanosecond) * time.Duration(100000*bits)
		dur2 := time.Duration(time.Nanosecond) * time.Duration(100000*timestampUpperBound)
		firstTime := epoch.Add(dur)
		secondTime := firstTime.Add(dur2)

		if secondTime.After(time.Now()) {
			return &firstTime
		}

		bits += timestampUpperBound
	}

	return nil
}

// power raises one int to the power of another, an operator oddly missing from golang
func (q *quoteNumberType) power(x int64, n int32) int64 {
	if n == 0 {
		return int64(1)
	}
	if n == 1 {
		return x
	}

	even := n%2 == 0
	xsq := x * x
	if even {
		newPow := int32(n / 2)
		return q.power(xsq, newPow)
	}
	newPow := int32((n - 1) / 2)
	return x * q.power(xsq, newPow)
}

// IsValid ...
func (q *quoteNumberType) IsValid(qn *string) bool {
	if qn == nil {
		return false
	}

	qns := *qn

	if len(qns) > maxLegalQuoteNumberLength {
		return false
	}

	if len(qns) < minLegalQuoteNumberLength {
		return false
	}

	return true
}

// ValidOrNew ...
func (q *quoteNumberType) ValidOrNew(qn *string) string {
	if qn == nil || !q.IsValid(qn) {
		return q.Create(nil)
	}

	return *qn
}

// createForIdNameDOB creates a 12-char identifier from a JobNumber, birthdate, and last name
// JobNumber and birthdate are recoverable; lastname serves as a confirmation hash
func (q *quoteNumberType) createForIdNameDOB(id int32, name string, y, m, d int) string {
	saltyBits := int64(q.idNameDOBToInt(id, name, y, m, d))
	chompArray := q.toChomps(saltyBits, 12)
	swirlyChomps := q.confusticate(chompArray)

	return q.chompArrayToChompTokens(swirlyChomps)
}

// parseIdDOBNameHash ...
func (q *quoteNumberType) parseIdDOBNameHash(quoteIdentifier *string) (int32, uint32, int, int, int) {
	if quoteIdentifier == nil || *quoteIdentifier == "" {
		return 0, 0, 0, 0, 0
	}
	cleanID := ""
	for _, thisChar := range *quoteIdentifier {
		uc := strings.ToUpper(string(thisChar))
		if _, found := idNumbers[uc]; found {
			cleanID += uc
		}
	}
	swirlyChomps := q.fromChompTokens(cleanID)
	chompArray := q.deconfusticate(swirlyChomps)
	saltyBits := uint64(q.toNumber(chompArray))

	return q.intToIDNameDOB(saltyBits)
}

func (q *quoteNumberType) idNameDOBToInt(id int32, name string, doby, dobm, dobd int) uint64 {
	rtn := uint64(0)

	// Use 12 bits from the name for the high-order bits
	nameInt := q.toSixteenBitHash(name) & 0xfff

	// make room for the other 48 bits
	rtn = rtn | (uint64(nameInt) << 48)

	// put all 32 bits of the passed-in id in the middle
	rtn = rtn | (uint64(id) << 16)

	// we'll keep 16 bits of days for the DOB -- that's almost 180 years
	dobInt := uint64(q.daysAfterBirthdateEpoch(doby, dobm, dobd)) & 0xffff

	rtn = rtn | dobInt

	return rtn
}

func (q *quoteNumberType) intToIDNameDOB(key uint64) (int32, uint32, int, int, int) {
	dobInt := int32(key & 0xffff)
	doby, dobm, dobd := q.intToBirthdate(dobInt)
	key = key >> 16
	id := int32(key & 0xffffffff)
	key = key >> 32

	return id, uint32(key), doby, dobm, dobd
}

// Guessing that as I write this we don't insure anyone over 106 years old
var birthdateEpoch = time.Date(1910, 1, 1, 12, 0, 0, 0, time.UTC)

func (q *quoteNumberType) daysAfterBirthdateEpoch(y, m, d int) int32 {
	dob := time.Date(y, time.Month(m), d, 12, 0, 0, 0, time.UTC)

	if !dob.After(birthdateEpoch) {
		return 0
	}
	afterEpochDur := dob.Sub(birthdateEpoch)
	return int32(afterEpochDur.Hours() / 24)
}

func (q *quoteNumberType) intToBirthdate(bdd int32) (int, int, int) {
	dur := time.Duration(bdd*24) * time.Hour
	rtn := birthdateEpoch.Add(dur)

	return rtn.Year(), int(rtn.Month()), rtn.Day()
}

// pad with 21 since it makes a nice 10101 pattern
const padByte = byte(21)
const fiveBits = byte(31)
const sixteenBits = uint32(65535)
const ignoreChars = byte(97)

func (q *quoteNumberType) toSixteenBitHash(name string) uint32 {
	asciiFinder := regexp.MustCompile("[a-z]")
	buf := strings.Join(asciiFinder.FindAllString(strings.ToLower(name), -1), "")
	if len(buf) == 0 {
		return uint32(0)
	}
	origBuf := buf
	lenSalt := uint32(len(buf) % 4)
	// need at least 6 characters to do anything useful
	for len(buf) < 6 {
		buf += origBuf
	}
	bufPos := 0
	bufLength := len(buf)
	hash := uint32(0)
	// Prime the frame with the length of the name
	currentFrameLength := uint(2)
	frame := lenSalt

	// Keep moving through the buffer until we get to the end
	for bufPos < bufLength-1 {
		// load the frame with more bits
		for currentFrameLength < 16 {
			bitsToLoad := padByte
			// don't start padding until we run out of characters
			if bufPos < bufLength {
				// we're working with lowercase utf-8 here, so we don't care about the first 96
				// subtract those 96 and take the 5 low-order bits
				bitsToLoad = (buf[bufPos] - ignoreChars) & fiveBits
				bufPos++
			} // end grab next byte
			newBitsMagnified := uint32(bitsToLoad) << currentFrameLength
			frame = newBitsMagnified | frame
			currentFrameLength += 5
		} // end load frame

		// _ = "breakpoint"

		// the frame is loaded, so now we pop the low-order 16 bits and XOR them with the old hash
		freshBits := frame & sixteenBits
		frame = frame >> 16
		currentFrameLength = currentFrameLength - 16
		// add the fresh bits to the hash
		hash = hash ^ freshBits
	}

	// _ = "breakpoint"
	// smash down to 12 bits
	hash = hash & 0xfff
	return hash
}
