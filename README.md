## clink
A distributed (see [Design](#design)) network monitoring tool that is easily operated from the command line. Clink operates via `icmp` ping requests while also collecting other statistics. Clink aims to provide a comprehensive monitoring view of services over a period of time. 

### <a name="why"></a>Why clink?
Clink was created out of a need for a more versital and long running CLI monitoring tool that provided detailed statistics while maintaining the simplicity of traditional 'query-once' network utilities. A key motivation for clink was to provide more comprehensive view of the data collected in a report format that is easily understood. 

#####The initial design goals of clink:
* Ability to gather scan information from multiple contexts.
 * External nodes powered by companion application `clinkd` provide different contexts.
 * Show statistics of such a scan in a meaningful manner. 
* Configuration files detailing monitoring configurations.
 * Simple human readable format. 
 * Self describing document structure.
 * Enable user to create comlpex scans that are easily reproducable.
* Suitable for persistent long running scans as well as shorter 'query-like' scans.
 * Provides detailed reports representing such data efficiently. 
 * Configuration to detail scan intervals, summary type, etc.

Clink is designed to provide more power in a simpler tool. 

### <a name="howitworks"></a>How It Works
Clink is accompanied...

### <a name="design"></a>Design
To follow... 

### <a name="usage"></a>Usage
Clink has several modes that can be used to accomplish a variety of tasks. The core functionality of clink provides an interface to monitor a services based on the specification in a config file. ... 

#####<a name="options"></a>Options

```
-s  	...
-f 		...
```

#####Example Usage

```bash
$ clink -s ...
```

```bash
$ clink -s ...
```

### <a name="build"></a>Clink Builds

#####Distributions
Visit the [builds](https://github.com/mbergoon/clink/releases) page for pre-built distributions.

#####Building from Source
Use the following command to retrieve the source: 
```
go get github.com/mbergoon/clink
```

Next, build the project with:
```
go install github.com/mbergoon/clink
```

The binary should be available in the `bin` folder of the Go workspace. Run it with (sudo is required to send icmp requests):
```
sudo clink [options]
```

### <a name="usecases"></a>Use Cases


### <a name="contributors"></a>Contributors
* cbergoon
* mbergoon

### <a name="lisense"></a>License
To follow...

