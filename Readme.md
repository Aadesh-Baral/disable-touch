# Ubuntu Touchscreen disabler
Ubuntu Touchscreen disabler helps to temporarily disable touchscreen in your Ubuntu device. This is written in golang. So you should have golang installed on your device.

## Dependencies
[Golang](https://go.dev/dl/)

### Install
- Make sure you have Golang installed on your system.
- Provide required permissions with:
    ```bash
    sudo chmod +x scripts/*.sh
    ```
- Add your touchsreen driver name in `driver_prefixs.txt`
  - To get your driver name enter `xinput --list` in your terminal, and look for your touchscreen driver. Dont be confused with your touchpad driver.
- Then run `go run main.go` to disable touchscreen.
