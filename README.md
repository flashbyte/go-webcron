[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![Go Version][goversion-shield]](https://go.dev)

<h3 align="center">go-webcron</h3>

  <p align="center">
    Tiny webcron build in go
    <br />
    <br />
    <a href="https://github.com/flashbyte/go-webcron/issues">Report Bug</a>
    ·
    <a href="https://github.com/flashbyte/go-webcron/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
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
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This application call in specified time interval a URL via get request. All it does is report if the endpoint was reachable. 

The need for this app came from using [nextcloud](https://nextcloud.com) which requiers an external trigger for maintenance tasks.

The app is meant to be deployed in a container and just do its job.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Built With

* ![Go Version](https://img.shields.io/github/go-mod/go-version/flashbyte/go-webcron?style=plastic)


<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started


### Prerequisites

* go

### Compile

```sh
go build cron.go
```


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage localy

```sh
export URL=http://google.com
export TIME=30
export DEBUG=true
./cron
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/flashbyte/go-webcron.svg?style=for-the-badge
[contributors-url]: https://github.com/flashbyte/go-webcron/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/flashbyte/go-webcron.svg?style=for-the-badge
[forks-url]: https://github.com/flashbyte/go-webcron/network/members
[stars-shield]: https://img.shields.io/github/stars/flashbyte/go-webcron.svg?style=for-the-badge
[stars-url]: https://github.com/flashbyte/go-webcron/stargazers
[issues-shield]: https://img.shields.io/github/issues/flashbyte/go-webcron.svg?style=for-the-badge
[issues-url]: https://github.com/flashbyte/go-webcron/issues
[license-shield]: https://img.shields.io/github/license/flashbyte/go-webcron?style=for-the-badge
[license-url]: https://github.com/flashbyte/go-webcron/blob/master/LICENSE.md
[goversion-shield]: https://img.shields.io/github/go-mod/go-version/flashbyte/go-webcron?style=for-the-badge
