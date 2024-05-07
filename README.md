**Project Name: DNS Health Checker**

## Description

This is a Golang project that implements a health checker website using DNS. The purpose of this project is to check the health status of a domain by performing DNS queries.

### Curl Command to Check Domain Health

To check the health of a domain, you can use the following curl command:

```
curl -X POST -H "Content-Type: application/json" -d '{"domain": "example.com"}' https://health-checker-website.onrender.com/check
```

Replace `"example.com"` with the domain you want to check.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
