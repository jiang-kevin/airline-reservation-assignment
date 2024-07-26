# airline-reservation-assignment

## Project Setup
1. Clone the repository
2. To build, run the following command. Dependencies will be installed automatically:
```
go build -o <OUTPUT_FILENAME>
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
