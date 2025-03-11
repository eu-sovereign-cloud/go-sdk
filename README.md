# go-sdk

Go SDK for the SECA API specification. This SDK is currently mostly handwritten but will be replaced using code generators for the majority of the code in the future. The client HTTP code is already generated using a code generator.

## Getting Started

To get started with the project, follow these steps:

1. Clone the repository:

    ```sh
    git clone git@github.com:eu-sovereign-cloud/go-sdk.git
    cd go-sdk
    ```

2. Initialize the submodule:

    ```sh
    git submodule init
    ```

3. Update all external dependencies:

    ```sh
    make clean update all mock
    ```

## Testing

To execute unit and integration tests, run the following command:

```sh
make test
```
