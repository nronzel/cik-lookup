# CIK Lookup

## Overview

This is a simple and efficient CLI written in Go to lookup the Central Index Key
(CIK) numbers of companies from the SEC database. Enter a company name, and it
will return the associated CIK number.

## Features

- Handles gzip-compressed responses
- Extracts CIK and company name from the response
- Parses multi-word company names

## Requirements

- Go v1.20
- goquery package (used for HTML parsing)

## Installation

Clone the repo:

```bash
git clone https://github.com/nronzel/cik-lookup.git
cd cik-lookup
```

Ensure you have goquery installed:

```bash
go get -u github.com/PuerkitoBio/goquery
```

## Usage

To use the tool, run the following command from the terminal:

```bash
go run main.go <company_name>
```

Replacing the company name with the company you want to look up.

Example:

```bash
go run main.go microsoft corp
```

### Note

The results are truncated to 100 if more than 100 results are found. Try being
more specific in these situations.

## How It Works

Simple POST request with the company name to the [SEC CIK lookup URL](https://www.sec.gov/edgar/searchedgar/cik).

## Contributing

If you want to contribute to this and expand upon its functionality, please fork
the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License.

## Contact

For any queries or suggestions, feel free to open an issue.
