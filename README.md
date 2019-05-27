# Gatfish

![alt-text](https://cdn.drawception.com/images/panels/2015/1-5/PChhFFR6ym-8.png)

## GATFISH IS STILL IN DEVELOPMENT ALL DOCUMENTATION STATING HOW TO DO THINGS IS SIMPLY BEING PROACTIVE WITH THE DOCUMENTATION

Gatfish is a CLI application that governs AWS resources for proper tagging. It will do this by taking a YAML file of desired tags per AWS resource and their level of requirement. Gatfish will be able to then perform actions based of these levels such as logging from highest to lowest severity, removing severe culprits, adding tags that can be automatically added.

### Current Stage: PLANNING
[TO DO](TODO.md)

## Installation

TODO: Simple binary installation. Want to look for a third party binary lib.

## Set Up

Gatfish uses a YAML file for configuring of the agent, more information on setting up the configuration here, **FUTURE LINK TO CONFIG SETTINGS PAGE** as of now the file **MUST** be named `gatfish.yaml` and reside in the same directory as the binary of the agent.

## How To Run

TODO: Should this be a CLI with flags? or simply an agent that runs based on the YAML file? I feel like this might want to follow the route of Narkwhal and come with the options of a silent mode. However with this I feel like the silent mode should be the default unlike Narkwhal.

## Authors

* **Alex Costa** 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details