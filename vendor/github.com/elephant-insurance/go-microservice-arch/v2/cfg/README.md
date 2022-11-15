# configuration
The Elephant configuration package, which handles app configuration loading, reconciliation, and validation at startup.

## NOTE: this README is out of date as of March 2021.


## Overview
This package is intended for handling complex configuration scenarios for Go microservices. Starting with a base (dev) configuration, which contains everything necessary to run the service in the dev environment, additional configuration settings can be merged to create configurations for other environments. These additional settings can come from environment-specific "override" config files, which only need to contain the specific settings that change when moving from the dev environment. They may also come from environment variables on the host machine.

The package also performs basic validation, ensuring that config settings not marked as "optional" are not empty after merging. The validation function will also check to ensure that any "tokenized" settings are replaced properly in the final config. More complex validation is easily added by implementing the PreValidated and PostValidated interfaces on your config class.


## Getting started
The easiest way to implement configuration in a new microservice is to copy an existing one. The ms-geo-data project is specifically intended to serve as a template for new microservices, so it's the best place to start. Add the config folder, config.go, and config.yml to your project just as they are in ms-geo-data. The application config namespace, defined by these files, contains a definition of the config structure for this app, an initialization function to call LoadConfig and Validate, and (optionally) implementations of PreValidation and PostValidation. It's is also a great place to put global constants or variables that must be accessible from multiple other namespaces ("controllers", "models", etc.). To avoid circular references, the app-specific config package should never import any other packages from the app.

Next you'll want to customize your config. Add any new fields that you need to the configuration class, and rename the configuration class to reflect your new app. Remember that any optional fields should be marked with the config:"optional" tag. Be sure to add those same fields to your config.yml file. The values in your base config file should be the values you need to run the app on your development machine. The RequiredConfig setting in the existing config cannot be removed, though of course you may change the values in RequiredConfig.


## Varying configuration in other environments
Your configuration will change in other environments. To configure the app for other environments, there are several options.

### Override config files
Currently (March 2019) all of our processes use simple override files. With this system, we create an override config file for each environment. Currently these override config files are all stored in the service_configuration GitHub repository. The override config files only need to contain values that need to change as the app moves from your dev box to the target environment. When the app is deployed to Docker, the deployment process selects the appropriate override file and mounts in the location specified by the OverrideConfigPath setting in the base config. The configuration package handles the override logic, producing a runtime configuration that combines the settings from the two files, preferring the override file whenever there is a conflict.

### Tokenized configuration
Another approach to configuration, the approach that we're moving toward, uses tokenized configuration files to automate the process of replacing values in different environments. To configure an app this way, we need to add a tokenized config file, named config-tokens.yml by default, which contains all of the fields that change across environments. Instead of specifying a value for each setting, we supply a unique token agreed upon with the Automation team. For example, if the setting was named DatabasePassword, our base config.yml would have a value that would allow us to have a database connection in dev. In the config-tokens.yml file, we might set the value to "SED_DATABASE_PASSWORD". Our automated deploy process will dynamically create the override config file from the tokenized config file by replacing each token with a real value, known by the Automation team and appropriate to the environment. In this way, a developer never needs to know (in this example) the database password for our production environment.

Note that the token-replacement process is handled entirely by our automated deployment process. In case this process fails for any reason, this configuration package will check for non-replaced tokens in the final config. If it finds that the value of a particular setting is the same as its value in the tokenized config file, the process will halt with an error message.

### Base64-encoded override file
Sometimes it is necessary to save the override configuration in a format that does not preserve the whitespace formatting required by YAML. To remedy this problem, the override file may be base-64 encoded and renamed to contain the string "-base64" (configurable by the public variable Base64OverrideToken) in its path. If the config package encounters an override file with "_base64" in its path, it will attempt to base-64 decode the override file before parsing it.

TODO:

replace override func with merge
remove pre-processing of tokens file for now
add Transformers for Logging, Diagnostics, RequiredConfig
think through environment vars -- override after merge
add token check to validation
