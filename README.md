# Clock Application

### Project Description
This Go application is a simple clock that prints specific values at set intervals:
- "tick" every second
- "tock" every minute
- "bong" every hour

The application is designed to run for three hours and then exit. The time it can run and exit can also be adjusted. Additionally, it allows users to alter the printed values at runtime without stopping the application.

### Installation Instructions

1. Download the clock-app repository / clone the clock-app repository.
    ```bash
    cd clock-app
    ```

### Option 1: Build and Run with Docker
You can build and run the application using Docker. A bash script **build_and_run.sh** is provided to simplify this process.

1. Ensure Docker is installed and running on your system. 
2. Login into Docker
    ```bash
    Docker login
    ```
3. Run the build_and_run.sh script
    ```bash
    ./build_and_run.sh
    ```
    
The script performs the following steps:

1. Checks if the Docker image clock-app already exists.
2.  If the image does not exist, it builds the Docker image.
3.  Runs the Docker container interactively.
    
### Option 2: Build and Run with Makefile
You can also build and run the application using Makefile.

1. Ensure Go is installed on your system.
2. Build the Application
    ```bash
    make build
    ```
3. Run the Application
   ```bash
   make run
   ```
4. Run the tests
    ```bash
    make test
    ```

5. Clean up
    ```bash
    make clean
    ```
This provides clear instructions for both Docker and Makefile build processes. Ensure you update the repository URL and any other specific details based on your actual setup.

### Set Up Environment Variables

Create a **.env** file in the root directory and set the default values for TICK, TOCK, BONG, and CLOCK_LIMIT. Example:

```bash
TICK="tick"
TOCK="tock"
BONG="bong"
CLOCK_LIMIT=3000 # The number of seconds the clock should run
```

### Usage Instructions

When the application is running, you can interact with it via the console to change the printed messages or stop the application.

###### Changing Messages

To change the `tick`, `tock`, and `bong` values or quit the application, follow these steps:

1. Disable print first by toggling the print settings.
2. Enter the necessary commands to change the messages.
3. Enable print by toggling the print settings again.

###### Commands
- `tick [value]`: Change the default "tick" message.
  - Example: `tick quack`
- `tock [value]`: Change the default "tock" message.
  - Example: `tock click`
- `bong [value]`: Change the default "bong" message.
  - Example: `bong gong`
- `t`: Toggle the print settings (enable/disable printing).
- `quit`: Stop the clock application.

##### Features
1. **Interval Printing:** Prints ***"tick"*** every second, ***"tock"*** every minute, and ***"bong"*** every hour.
2. **Runtime Configuration:** Allows dynamic changes to the printed values while the application is running.
3. **Graceful Shutdown:** Handles graceful shutdown on receiving ***SIGINT*** or ***SIGTERM*** signals.
4. **Configurable Run Duration:** The clock runs for a configurable duration (default is three hours) and then exits automatically.


### Test Files

To run all the test files either run 

```bash
go test -v ./...
```

or

```bash
make test
```

##### main_test.go

This file contains unit tests for the **`handleUserCommand`** function to ensure it processes user commands correctly. It tests various scenarios, including setting the tick, tock, and bong values, as well as handling unknown commands.

##### config_test.go

This file includes tests for the configuration package. It verifies the functionality of environment variable retrieval functions and ensures that the **`LoadEnv`** function correctly loads and parses environment variables for the application's configuration.

##### clock_test.go

This file contains comprehensive tests for the `Clock` struct and its methods. The tests include:
- **TestNewClock**: Verifies the creation of a new `Clock` instance and checks its initial values.
- **TestClock_Run**: Ensures the `Run` method operates correctly, including starting and stopping the clock.
- **TestClock_TogglePrint**: Verifies the `TogglePrint` method toggles the print state correctly.
- **TestClock_SetValues**: Ensures that the `SetTick`, `SetTock`, and `SetBong` methods update the clock's values correctly.
- **TestClock_PrintFunctions**: Verifies that the print functions (`printTick`, `printTock`, `printBong`) operate correctly based on the print state.
- **TestClock_StopTicker**: Tests the closing of the `stopTicker` channel.
- **TestClock_Finished**: Tests the closing of the `Finished` channel when the clock's time limit is reached.

These tests ensure the reliability and correctness of the core functionality of the clock application.








