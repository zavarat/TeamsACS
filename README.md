# TeamsACS

TeamsACS is committed to providing work teams with exceptional ease of network management. We use Mikrotik's network products as the core foundation, while extending the system's capabilities to a wider range of network devices, such as OpenWrt.

The core of the system is based on Golang technology, providing excellent performance and ease of deployment.

## Systems Architecture

![image](https://user-images.githubusercontent.com/377938/97301570-e28b3d80-1892-11eb-85a8-5cc5f80449a4.png)

## System Features

### TR069 ACS integration

Preferring GenieACS open source ACS system integration， GenieACS can work with any device that supports the TR-069 protocol.

It auto-discovers the device’s parameter tree (including vendor-specific parameters) making no assumptions about the device’s data model.

It’s been tested with a wide range of devices (DSL, cable, fiber optics, LTE CPEs, VoIP phones) from many different manufacturers.

It is also the officially recommended system by Mikrotik, which has been tested extensively and is safe to use.

### Northbound Interface

- Provides a unified API for various third-party management systems, based on the HTTPS Json protocol.
