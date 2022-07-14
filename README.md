# Log Toolkit

Es una libreria de registro de logs que implementa zap(uber), soporta niveles de logs. 
Recibe como parametro un mensaje y opcionalmente la estructura a la cuál se le 
desea dar un seguimiento.
Desarrollada en go. 

## Como utilizar:
Importar en archivo de fuentes de go:

`import log_toolkit "github.com/validatecl/log-toolkit"`

El primer parámetro es opcional

```golang
		logger := log_toolkit.NewLogToolkit(nil, &log_toolkit.ZapConfigInput{ 
			EnableConsole:     true,
			ConsoleLevel:      "debug",
			ConsoleJSONFormat: true,
		})
```

En el caso de querer acceder a los mensajes o hacer testings:

1.- Importar las librerías "go.uber.org/zap/zapcore" y "go.uber.org/zap/zaptest/observer"

2.- Crear una instancia de la librería observer pasando como parámetro el nivel de logs requerido.
Esta nos retorna dos objetos de tipo zap.core y observedLogs respectivamente.

3.- Crear instancia de zap (logger) usando como parámetro la variable core.

4.- Ahora si, nuestro logger lo inicializamos con los valores 


```golang
		core, recorded := observer.New(zapcore.DebugLevel)
		zap_core_logger := zap.New(core)
		logger := log_toolkit.NewLogToolkit(zap_core_logger, &log_toolkit.ZapConfigInput{
			EnableConsole:     true,
			ConsoleLevel:      "debug",
			ConsoleJSONFormat: true,
        })
```


        
Los niveles de logs disponibles en log-toolkit son:

    // DebugLevel logs are typically voluminous, and are usually disabled in
    // production.
    DebugLevel Level = iota - 1
    // InfoLevel is the default logging priority.
    InfoLevel
    // WarnLevel logs are more important than Info, but don't need individual
    // human review.
    WarnLevel
    // ErrorLevel logs are high-priority. If an application is running smoothly,
    // it shouldn't generate any error-level logs.
    ErrorLevel
    // PanicLevel logs a message, then panics.
    PanicLevel
    // FatalLevel logs a message, then calls os.Exit(1).
    FatalLevel

Nota: Es importante considerar que al utilizar el FatalLevel no es aplicable la función recover()
de GO.

Para más információn consultar la documentación de ZapCore https://godoc.org/go.uber.org/zap/zapcore

Los parámetros de entrada son un mensaje de error y una structura x de errores.

```golang
		logger.Debug("Debugging", interface{})
		logger.Info("Mensaje informativo")
		logger.Warn("Mensaje de advertencia")
        logger.Error("Mensaje de error") 
        logger.Panic("Error grave, el sistema puede caerse sino se maneja la excepción") 
        logger.Fatal("Error irrecuperable, se cerrará la aplicación") 
```

Para acceder al stdout y recorrer los mensajes

```golang
	for _, logs := range recorded.All() {
        fmt.Println(logs.Message)
        fmt.Println(logs.Level)
    }
```


## Como hacer test del proyecto:

Para hacer test del proyecto simplemente se necesita ejectuar el comando `make test`, esto genera una salida similar a esta:
```
Ejecutando tests...
=== RUN   TestBaseConfigurer
=== RUN   TestBaseConfigurer/Default_configuration
2020/01/08 11:52:41 default Filename value: 08-01-2020.log
2020/01/08 11:52:41 default MaxSize value: 100 MB 
2020/01/08 11:52:41 default TimeFormat value: 2006-02-01 15:04:05.000 
2020/01/08 11:52:41 default FilePath value: logs 
=== RUN   TestBaseConfigurer/Custom_configuration
2020/01/08 11:52:41 default TimeFormat value: 2006-02-01 15:04:05.000 
2020/01/08 11:52:41 default FilePath value: logs 
=== RUN   TestBaseConfigurer/Custom_configuration#01
2020/01/08 11:52:41 default Filename value: 08-01-2020.log
2020/01/08 11:52:41 default MaxSize value: 100 MB 
2020/01/08 11:52:41 default TimeFormat value: 2006-02-01 15:04:05.000 
2020/01/08 11:52:41 default FilePath value: logs 
=== RUN   TestBaseConfigurer/Custom_configuration_2
2020/01/08 11:52:41 default Filename value: 08-01-2020.log
2020/01/08 11:52:41 default MaxSize value: 100 MB 
2020/01/08 11:52:41 default TimeFormat value: 2006-02-01 15:04:05.000 
2020/01/08 11:52:41 default FilePath value: logs 
=== RUN   TestBaseConfigurer/Custom_configuration_3
2020/01/08 11:52:41 default Filename value: 08-01-2020.log
2020/01/08 11:52:41 default MaxSize value: 100 MB 
2020/01/08 11:52:41 default TimeFormat value: 2006-02-01 15:04:05.000 
2020/01/08 11:52:41 default FilePath value: logs 
--- PASS: TestBaseConfigurer (0.00s)
    --- PASS: TestBaseConfigurer/Default_configuration (0.00s)
    --- PASS: TestBaseConfigurer/Custom_configuration (0.00s)
    --- PASS: TestBaseConfigurer/Custom_configuration#01 (0.00s)
    --- PASS: TestBaseConfigurer/Custom_configuration_2 (0.00s)
    --- PASS: TestBaseConfigurer/Custom_configuration_3 (0.00s)
=== RUN   TestMarshallSimpleFields
--- PASS: TestMarshallSimpleFields (0.00s)
=== RUN   TestMarshalFields
=== RUN   TestMarshalFields/Verificando_cantidad_de_registros.(Success)
Result: [{example 15 0 valueexample1 valueexample1} {child 15 0 valuechild valuechild} {child 15 0 valuechild3 valuechild3} {child 15 0 valuechild4 valuechild4} {child 15 0 valueptrChild valueptrChild} {child 15 0 valuechild3 valuechild3} {child 15 0 valuechild4 valuechild4} {map1 15 0 map_1 map_1} {map2 15 0 map_2 map_2}]--- PASS: TestMarshalFields (0.00s)
    --- PASS: TestMarshalFields/Verificando_cantidad_de_registros.(Success) (0.00s)
=== RUN   TestLogToolkit
=== RUN   TestLogToolkit/NewLogToolkit_creation
=== RUN   TestLogToolkit/Info_Log_default_logger
2020/01/08 11:52:41 default Filename value: 08-01-2020.log
2020/01/08 11:52:41 default MaxSize value: 100 MB 
2020/01/08 11:52:41 default TimeFormat value: 2006-02-01 15:04:05.000 
2020/01/08 11:52:41 default FilePath value: logs 
{"level":"info","date":"2020-08-01 11:52:41.322","caller":"testing/testing.go:909","payload":"Prueba Testing Info: Hola Mundo"}
{"level":"warn","date":"2020-08-01 11:52:41.322","caller":"testing/testing.go:909","payload":"Prueba Testing Warn: Hola Mundo"}
{"level":"error","date":"2020-08-01 11:52:41.322","caller":"testing/testing.go:909","payload":"Prueba Testing Error: Hola Mundo"}
=== RUN   TestLogToolkit/Debug_Log_default_logger
2020/01/08 11:52:41 default Filename value: 08-01-2020.log
2020/01/08 11:52:41 default MaxSize value: 100 MB 
2020/01/08 11:52:41 default TimeFormat value: 2006-02-01 15:04:05.000 
2020/01/08 11:52:41 default FilePath value: logs 
{"level":"debug","date":"2020-08-01 11:52:41.322","caller":"testing/testing.go:909","payload":"Prueba Testing Debug: Hola Mundo"}
{"level":"info","date":"2020-08-01 11:52:41.322","caller":"testing/testing.go:909","payload":"Prueba Testing Info: Hola Mundo"}
{"level":"warn","date":"2020-08-01 11:52:41.322","caller":"testing/testing.go:909","payload":"Prueba Testing Warn: Hola Mundo"}
{"level":"error","date":"2020-08-01 11:52:41.322","caller":"testing/testing.go:909","payload":"Prueba Testing Error: Hola Mundo"}
=== RUN   TestLogToolkit/Warn_Log_default_logger
2020/01/08 11:52:41 default Filename value: 08-01-2020.log
2020/01/08 11:52:41 default MaxSize value: 100 MB 
2020/01/08 11:52:41 default TimeFormat value: 2006-02-01 15:04:05.000 
2020/01/08 11:52:41 default FilePath value: logs 
{"level":"warn","date":"2020-08-01 11:52:41.322","caller":"testing/testing.go:909","payload":"Prueba Testing Warn: Hola Mundo"}
{"level":"error","date":"2020-08-01 11:52:41.322","caller":"testing/testing.go:909","payload":"Prueba Testing Error: Hola Mundo"}
=== RUN   TestLogToolkit/Error_Log_default_logger
2020/01/08 11:52:41 default Filename value: 08-01-2020.log
2020/01/08 11:52:41 default MaxSize value: 100 MB 
2020/01/08 11:52:41 default TimeFormat value: 2006-02-01 15:04:05.000 
2020/01/08 11:52:41 default FilePath value: logs 
{"level":"error","date":"2020-08-01 11:52:41.322","caller":"testing/testing.go:909","payload":"Prueba Testing Error: Hola Mundo"}
=== RUN   TestLogToolkit/Debug_Log
=== RUN   TestLogToolkit/Info_Log
=== RUN   TestLogToolkit/Warn_Log
=== RUN   TestLogToolkit/Error_Log
=== RUN   TestLogToolkit/Panic_Log
Prueba Testing Panic: Hola Mundo--- PASS: TestLogToolkit (0.00s)
    --- PASS: TestLogToolkit/NewLogToolkit_creation (0.00s)
    --- PASS: TestLogToolkit/Info_Log_default_logger (0.00s)
    --- PASS: TestLogToolkit/Debug_Log_default_logger (0.00s)
    --- PASS: TestLogToolkit/Warn_Log_default_logger (0.00s)
    --- PASS: TestLogToolkit/Error_Log_default_logger (0.00s)
    --- PASS: TestLogToolkit/Debug_Log (0.00s)
    --- PASS: TestLogToolkit/Info_Log (0.00s)
    --- PASS: TestLogToolkit/Warn_Log (0.00s)
    --- PASS: TestLogToolkit/Error_Log (0.00s)
    --- PASS: TestLogToolkit/Panic_Log (0.00s)
=== RUN   TestOsInformation
=== RUN   TestOsInformation/Valida_que_el_hostname_no_sea_Nulo_y_que_sea_de_tipo_string
=== RUN   TestOsInformation/Valida_que_el_hostname_no_sea_Nulo_y_que_sea_de_tipo_Int
--- PASS: TestOsInformation (0.00s)
    --- PASS: TestOsInformation/Valida_que_el_hostname_no_sea_Nulo_y_que_sea_de_tipo_string (0.00s)
    --- PASS: TestOsInformation/Valida_que_el_hostname_no_sea_Nulo_y_que_sea_de_tipo_Int (0.00s)
PASS
ok      github.com/validatecl/log-toolkit    0.004s
```

Del mismo modo, se puede utilizar el comando `make coverage`, esto genera una salida similar a esto:
```
Coverage...
github.com/validatecl/log-toolkit/basic_zap_config.go:66:    NewBaseConfig           100.0%
github.com/validatecl/log-toolkit/basic_zap_config.go:70:    GenerateConfig          100.0%
github.com/validatecl/log-toolkit/basic_zap_config.go:74:    newZapCore              100.0%
github.com/validatecl/log-toolkit/basic_zap_config.go:107:   validateZapConfig       100.0%
github.com/validatecl/log-toolkit/basic_zap_config.go:127:   getZapLevel             100.0%
github.com/validatecl/log-toolkit/basic_zap_config.go:144:   getEncoder              100.0%
github.com/validatecl/log-toolkit/basic_zap_config.go:156:   customTimeEncoder       100.0%
github.com/validatecl/log-toolkit/logger.go:39:              NewLogToolkit           100.0%
github.com/validatecl/log-toolkit/logger.go:53:              Debug                   100.0%
github.com/validatecl/log-toolkit/logger.go:58:              Info                    100.0%
github.com/validatecl/log-toolkit/logger.go:63:              Warn                    100.0%
github.com/validatecl/log-toolkit/logger.go:68:              Error                   100.0%
github.com/validatecl/log-toolkit/logger.go:73:              Panic                   100.0%
github.com/validatecl/log-toolkit/logger.go:78:              Fatal                   0.0%
github.com/validatecl/log-toolkit/logger_field.go:20:        NewFieldMarshaller      100.0%
github.com/validatecl/log-toolkit/logger_field.go:25:        MarshalFields           100.0%
github.com/validatecl/log-toolkit/logger_field.go:35:        marshallRecursive       100.0%
github.com/validatecl/log-toolkit/os_information.go:12:      NewOsInformation        100.0%
total:                                                                                          (statements)            98.7%
```

##Como colaborar

Para colaborar con nuestra libreria es necesario creare un pull/merge request de una rama **feature** a la rama **develop** , es necesario que el request cumpla con ciertos guidelines los cuales se encuentran [aqui](https://ww/771818065/Guidelines+generales) 