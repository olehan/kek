<p align="center">
    <a href="https://github.com/olehan/kek">
        <img alt="Kek Logo" src="https://raw.githubusercontent.com/olehan/logo/master/kek.png" width="546">
    </a>
</p>

<p align="center">
    Zero allocated, structured, leveled and pretty Golang logger
</p>

<p align="center">
    <a href="https://travis-ci.org/olehan/kek">
        <img alt="Travis Status" src="https://travis-ci.org/olehan/kek.svg">
    </a>
    <a href="https://codecov.io/gh/olehan/kek">
        <img alt="Kek Code Coverage" src="https://codecov.io/gh/olehan/kek/branch/master/graph/badge.svg">
    </a>
    <a href="https://github.com/olehan/kek/blob/master/LICENSE">
        <img alt="Kek License" src="https://img.shields.io/github/license/mashape/apistatus.svg">
    </a>
    <a href="https://github.com/olehan/kek/releases">
        <img alt="Kek Last Release" src="https://img.shields.io/github/tag/olehan/kek.svg?label=release">
    </a>
</p>

----

<p align="center">
    <strong><a href="#intro">Intro</a></strong>
    |
    <strong><a href="#benchmarks">Benchmarks</a></strong>
    |
    <strong><a href="#usage">Usage</a></strong>
    |
    <strong><a href="#installation">Installation</a></strong>
    |
    <strong><a href="#license">License</a></strong>
</p>

----

<h2 id="intro">🌚 Intro</h2>

A lot of people don't care about the context stuff and expect to see the code,
output and specs, so without further a due:

***Code:***
```go
// Here we are, just creating our new logger with a default 'formatter'.
// We'll talk about formatters letter on Usage section...
logger := kek.NewLogger("logger", "name")
logger.Debug.Println("And that is it!")
```
***Output:***
```
logger.name | debug  2019/7/28 20:54:40.683859000 [4666]:   And that is it!
```

<h2 id="benchmarks">🏃 Benchmarks</h2>
Coming soon...

<h2 id="usage">🔧 Usage</h2>

Aight. If you're staying with me to see more complex usage of this lib - go into the [examples](examples) folder or keep reading.

#### Factories
These are things that help you to populate logger with a specific configuration that
is going to be 'linked' to a logger or just copied, so modification to the factory won't
affect on other loggers.

```go
factory := kek.NewFactory(
    // First arg is an io.Writer
    os.Stdout,
    // And the second is the formatter. Don't worry, we've got some
    // formatter packed up within the lib. But if you want, you
    // can create your formatter and give your logs
    // a new life.
    sugared.Formatter,
)

// Now let's modify this factory a little bit.
factory.
    SetWithColors(true).
    SetWithDate(false).
    SetWithMutex(true)

// Alright, now our logger have extended factory configurations.
logger := factory.NewLogger()
// And even can modify his own config, not affecting the factory.
logger.SetWithDate(true)

// But for those who wants to link loggers coming out of your
// factory, we've got NewLinkedLogger func.
// It returns the same logger, except its configurations are linked
// to the factory.
linkedLogger := factory.NewLinkedLogger()
// So any change made for the factory is going to affect a linked
// logger.
factory.SetWithNS(false)
```

#### Loggers
The logger is a thing that you use to write logs into a writer.
To use them you need to create a new one from the factory to extend
already specified configs or just use default configs by using the
public `NewLogger(name ...string) *Logger` function.
```go
key.NewLogger().Info.Println("yeah")
```

#### Formatters
Oke, now I know that there is a lot of projects that need a specific
log format and for that purposes to log stuff, you'll need
a formatter that will control every log. There is a set of configurable
(or not) default formatters made to save your time.

***Code:***
```go
minified := kek.NewLogger("minified").SetFormatter(minified.Formatter)
sugared := kek.NewLogger("sugared").SetFormatter(sugared.Formatter)
minified.Debug.Println("yeah")
sugared.Info.Println("YEAAAH")
``` 
***Output:***
```
minified - debug: yeah
sugared  | info  2019/7/28 21:27:27.788431000 [4715]: YEAAAH
```

#### Logger names
Every logger can have a name and a factory can also have a prefix for
a logger. This functionality is not necessary, but I should say that
is pretty useful when you need to separate loggers to get the context
of a process.

***Code:***
```go
factory := kek.NewFactory("service")
// Oh, and there is also a name tabulation, but it will take affect
// depends on your formatter.
factory.SetNameTabulation(true)

userPostsLogger := factory.NewLogger("user", "posts")
userCommentsLogger := factory.NewLogger("user", "comments")

// If a property 'WithColors' is enabled, you can set a random color
// for the loggers name.
userPostsLogger.SetRandomNameColor()

userPostsLogger.Debug.Println("Post edited")
userCommentsLogger.Warn.Println("Comments are refreshed")
```

***Output:***
```
# Welp, to see the colors - head over to the examples folder.
service/user.posts    | debug 2019/7/28 21:27:27.788431000 [4715]: Post edited
service/user.comments | warn  2019/7/28 21:27:27.788521000 [4715]: Comments are refreshed
```

#### Printer Types
The printer is the function that handles the job of executing a formatter and pool (🚰) management.
For a moment of this documentation kek delivers 3 types of printers:
* ***Base*** - the only printer that you've seen so far (`Print`, `Println`)
* ***Template*** - a printer that takes a format of a log and inserts values into it
(`PrintT`, `PrintTM`, `PrintTKV`)
* ***Structured*** - a printer that takes a message and writes key-value pairs in a specific structured way (`PrintSKV`)

***Code:***
```go
logger := kek.NewLogger()
logger.Debug.
    PrintT("wassup, {} - {}", "boi", 123).
    PrintTM("wassup, NewMap{{name}}, Age: {{age}}",
        ds.NewMap().
            Set("name", "Boi").
            Set("age", 17),
    ).
    PrintTM("wassup, Map{{name}}", ds.Map{
        "name": "Boi",
    }).
    PrintTKV("wassup, KeyValue{{name}} {{34}} {{134}} - {{unknown}}",
        "name", "Boi",
        "34", true,
        // Skips the key that is not a string.
        134, "ooh",
        // Skips the key value pair without a value.
        "unknown",
    )

logger.Info.
    PrintSKV("message for key values",
        "key1", "value1",
        "key2", 242456246,
        "key3", true,
        "key4", 135135.13413,
    )
```

***Output:***
```
debug 2019/7/28 21:53:20.501202000 [4775]:   wassup, boi - 123
debug 2019/7/28 21:53:20.501296000 [4775]:   wassup, NewMapBoi, Age: 17
debug 2019/7/28 21:53:20.501300000 [4775]:   wassup, MapBoi
debug 2019/7/28 21:53:20.501305000 [4775]:   wassup, KeyValueBoi true {{134}} - {{unknown}}
info  2019/7/28 21:56:20.501355000 [4811]:   message for key values
 > key1: value1
 > key2: 242456246
 > key3: true
 > key4: 135135.13413
```

<h2 id="installation">⬇️ Installation</h2>

```go
go get github.com/olehan/kek
```

<h2 id="license">🔖 License</h2>

It's [MIT](LICENSE).
What else would you expect? 🌚
