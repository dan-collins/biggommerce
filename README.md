<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
***
***
***
*** To avoid retyping too much info. Do a search and replace for the following:
*** dan-collins, biggommerce, twitter_handle, email, Big Gommerce, This is a library for consuming the Big Commerce API in go. Handles things like auth and [un]marshalling to hopefully make things easier.
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/dan-collins/biggommerce">
    <img src="images/logo80.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Big Gommerce</h3>

  <p align="center">
    This is a library for consuming the Big Commerce API in go. Handles things like auth and [un]marshalling to hopefully make things easier.
    <br />
    <a href="https://pkg.go.dev/github.com/dan-collins/biggommerce#section-directories"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/dan-collins/biggommerce/issues">Report Bug</a>
    ·
    <a href="https://github.com/dan-collins/biggommerce/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

If you want to interact with the BigCommerce rest API in a golang project, this library should
help you out. The plan is to eventually build it out for all endpoints/resources but for now most of
the [orders](https://developer.bigcommerce.com/api-reference/store-management/orders/) endpoint is addressed.

Module is still considered "unstable" (v0.x.x) until I have a chance to use this in a production app.

### Built With

* [Go](https://golang.org/)



<!-- GETTING STARTED -->
## Getting Started

To use the module in your project follow these simple steps.

### Prerequisites

[Go >= 1.15](https://golang.org/dl/) - Most versions should work, but was tested on 1.15.

### Installation

1. Go get the module
   ```sh
   go get github.com/dan-collins/biggommerce
   ```
2. Import the relevant packages
   ```go
   package main

   import "github.com/dan-collins/biggommerce/order"
   ```



<!-- USAGE EXAMPLES -->
## Usage

```go
package main

import (
	"fmt"
	// Make sure you import the package you want to access
	bgOrder "github.com/dan-collins/biggommerce/order"
	"time"
)

func main() {
	orders := GetOrders()
	fmt.Println(orders)
}

func GetOrders() []bgOrder.Order {
	// Fill in your BigCommerce details here
	bcToken := "{YOUR-TOKEN}"
	bgClient := "{YOUR-CLIENT-ID}"
	bcStoreKey := "{YOUR-STORE-KEY}"

	// Create the client
	client := bgOrder.NewClient(bcToken, bgClient, bcStoreKey)

	// Set up your query criteria 
	// (see https://pkg.go.dev/github.com/dan-collins/biggommerce@v0.2.0/order#OrderQuery for details)
	ny, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	// Just getting a date range for beginning and end of a day
	sDate := time.Date(2021, 3, 3, 0, 0, 0, 0, ny)
	eDate := time.Date(2021, 3, 3, 23, 59, 59, 0, ny)
	// I only want orders with status id of 8
	status := 8

	// Setup the query to use based on the criteria
	orderQuery := bgOrder.Query{StatusID: status, MinDateModified: sDate, MaxDateModified: eDate}
	
	// Get back the data you care about, Hydrated will include all resource linked objects on each order returned
	orders, err := client.GetHydratedOrders(orderQuery)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return *orders
}
```

_For method documentation, please refer to the [GoDoc](https://pkg.go.dev/github.com/dan-collins/biggommerce#section-directories)_



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/dan-collins/biggommerce/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Project Link: [https://github.com/dan-collins/biggommerce](https://github.com/dan-collins/biggommerce)

[Contact](https://www.danjcollins.com/#contact) Me Directly


<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [BigCommerce](https://www.bigcommerce.com)





<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/dan-collins/biggommerce.svg?logo=github&style=for-the-badge
[contributors-url]: https://github.com/dan-collins/biggommerce/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/dan-collins/biggommerce.svg?logo=github&style=for-the-badge
[forks-url]: https://github.com/dan-collins/biggommerce/network/members
[stars-shield]: https://img.shields.io/github/stars/dan-collins/biggommerce.svg?logo=github&style=for-the-badge
[stars-url]: https://github.com/dan-collins/biggommerce/stargazers
[issues-shield]: https://img.shields.io/github/issues/dan-collins/biggommerce.svg?logo=github&style=for-the-badge
[issues-url]: https://github.com/dan-collins/biggommerce/issues
[license-shield]: https://img.shields.io/github/license/dan-collins/biggommerce.svg?style=for-the-badge
[license-url]: https://github.com/dan-collins/biggommerce/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/dan-collins-8805b111/