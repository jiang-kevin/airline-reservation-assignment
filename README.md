# airline-reservation-assignment

## Project Setup
1. Clone the repository
2. To build, run the following command. Dependencies will be installed automatically:
```
go build -o <OUTPUT_FILENAME>
```

To run tests:
```
go test ./...
```

## Running the CLI
In the following examples, the executable name is `ara`. If you built the executable with a different name, replace it as needed.

Usage:
```
ara <action> <seat> <number of consecutive seats>
```

Book a seat:
```
ara BOOK A0 1
```

Cancel a reservation on a seat:
```
ara CANCEL A0 1
```

Book multiple seats:
```
ara BOOK A0 3
```

## Design Notes
* If you want to see a visual representation of the booked seats, check `data.txt`. `o` means unbooked, `x` means booked.
* The data file that reserved seats are stored in can be configured by editing `config.yml`
* The logic for handling consecutive seats is simple: starting from the input seat, consecutively book seats to the right until the number specified is reached.
    * If it encounters any already booked seats on the way, the whole operation fails
    * If it reaches the end of the row, it will fail
    * It does not respect aisles - seats should be booked starting from the leftmost seat if the same aisle is desired
    * The main motivation behind this logic is to keep it consistent between booking and cancelling seats. If the logic becomes complex, it would mean that the same cancel operation could have different results based on what is already booked. I think it would require some form of differentiating between customers booking the seats to be able to handle more complex logic.
* `cobra` and `viper` libraries were used to simplify CLI development and offer a good starting point for the application, while also providing many features
