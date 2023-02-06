package jwt_validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyToken(t *testing.T) {
	tk := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJfQnp2QnNfRm1wRmY1OTI5d09NcTVva3RxNXE3ZTJfdDdENWxubHIyVlFvIn0.eyJleHAiOjE2NzU1OTM2MTMsImlhdCI6MTY3NTU5MTgxMywianRpIjoiN2FiNzNlMDctNWNkNS00Mzc3LWE1ZDktODI1MmY3MDM3NGE1IiwiaXNzIjoiaHR0cHM6Ly9rZXljbG9hay1zaXQucGxhdGZvcm0ubnAuZG9jdG9yYW55d2hlcmUuY29tL3JlYWxtcy9EQUIyQiIsImF1ZCI6ImIyYi5hZG1pbi5wb3J0YWwiLCJzdWIiOiJiNDRkZmMwNi00MWY4LTQ5MDUtODNlOC00NTdmZGYyNDcyM2UiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJiMmIuYWRtaW4ucG9ydGFsIiwic2Vzc2lvbl9zdGF0ZSI6ImJhNWUwMjA3LTZiMmQtNDhmOS05OTQzLWY0Nzc1MmIxMTdhZiIsImFjciI6IjEiLCJyZXNvdXJjZV9hY2Nlc3MiOnsiYjJiLmFkbWluLnBvcnRhbCI6eyJyb2xlcyI6WyJyb2xlLnN1cGVyYWRtaW4iXX19LCJhdXRob3JpemF0aW9uIjp7InBlcm1pc3Npb25zIjpbeyJzY29wZXMiOlsiZmluYW5jZS53cml0ZWZpbmFuY2UiLCJmaW5hbmNlLmRlbGV0ZWZpbmFuY2UiLCJmaW5hbmNlLnJlYWRmaW5hbmNlIl0sInJzaWQiOiIyYmUxOTc3Mi0yM2UzLTQwYTUtOGM5Mi03NGQyOWQ3NGU3MTgiLCJyc25hbWUiOiJmaW5hbmNlc2VydmljZSJ9LHsic2NvcGVzIjpbInByb3ZpZGVyLnJlYWRkb2N0b3IiLCJwcm92aWRlci53cml0ZWhvc3BpdGFsIiwicHJvdmlkZXIucmVhZGhvc3BpdGFsIiwicHJvdmlkZXIud3JpdGVjbGluaWMiLCJwcm92aWRlci53cml0ZWRvY3RvciIsInByb3ZpZGVyLmRlbGV0ZWhvc3BpdGFsIiwicHJvdmlkZXIuZGVsZXRlY2xpbmljIiwicHJvdmlkZXIuZGVsZXRlZG9jdG9yIiwicHJvdmlkZXIucmVhZGNsaW5pYyJdLCJyc2lkIjoiZjhlYzA3NDUtMjdjOS00MGZhLWEwMTktODY3NDVmNmMwOGE5IiwicnNuYW1lIjoicHJvdmlkZXJzZXJ2aWNlIn0seyJzY29wZXMiOlsibm90aWZpY2F0aW9uLmRlbGV0ZW5vdGlmaWNhdGlvbiIsIm5vdGlmaWNhdGlvbi5yZWFkbm90aWZpY2F0aW9uIiwibm90aWZpY2F0aW9uLndyaXRlbm90aWZpY2F0aW9uIl0sInJzaWQiOiJmYTExNTU4OS05ZmQyLTRmOWEtOGYzNS1jZDAzMDZlZWFmZDciLCJyc25hbWUiOiJub3RpZmljYXRpb25zZXJ2aWNlIn0seyJzY29wZXMiOlsiaW52ZW50b3J5LmRlbGV0ZWludmVudG9yeSIsImludmVudG9yeS53cml0ZWludmVudG9yeSIsImludmVudG9yeS5yZWFkaW52ZW50b3J5Il0sInJzaWQiOiIzN2EwZTNmOS00YTdlLTQ4YWEtYWI1My1kZWNjZGI2ZDQ0ZDciLCJyc25hbWUiOiJpbnZlbnRvcnlzZXJ2aWNlIn0seyJzY29wZXMiOlsicG9saWN5LmRlbGV0ZXBvbGljeSIsInBvbGljeS5yZWFkcG9saWN5IiwicG9saWN5LndyaXRlcG9saWN5Il0sInJzaWQiOiJhZGU2M2IyNy1iYzczLTRhMWEtOTk4OC02ODQ4MGYyZmYwNzMiLCJyc25hbWUiOiJwb2xpY3lzZXJ2aWNlIn0seyJzY29wZXMiOlsiY2xhaW0ucmVhZGNsYWltIiwiY2xhaW0ud3JpdGVjbGFpbSIsImNsYWltLmRlbGV0ZWNsYWltIl0sInJzaWQiOiIyMWM0YWQ3Ny00MzZhLTQzYmEtOTg3Zi0zNDhjODE2OTYyNTAiLCJyc25hbWUiOiJjbGFpbXNlcnZpY2UifSx7InNjb3BlcyI6WyJsb2dpbiJdLCJyc2lkIjoiYTU5ODhmNWMtZGE0Zi00NmVmLWJmZDItMDI1Mzc5OWU2ZmJiIiwicnNuYW1lIjoiYWRtaW5wb3J0YWwifSx7InNjb3BlcyI6WyJtZW1iZXJzaGlwLnJlYWRtZW1iZXIiLCJtZW1iZXJzaGlwLndyaXRlbWVtYmVyIiwibWVtYmVyc2hpcC5kZWxldGVtZW1iZXIiXSwicnNpZCI6ImI4ZjMxNTAwLWUyOWEtNGZlOS05N2NkLThmMGI4ODU0MzJkZSIsInJzbmFtZSI6Im1lbWJlcnNoaXBzZXJ2aWNlIn0seyJzY29wZXMiOlsiaWFtLnJlYWR1c2VyIiwiaWFtLmRlbGV0ZXVzZXIiLCJpYW0ud3JpdGV1c2VyIl0sInJzaWQiOiIzZjVhNDM5NC00ZTRhLTQ0NTgtYmNmMS1jYjJjYjZlODk4OTIiLCJyc25hbWUiOiJpYW1zZXJ2aWNlIn0seyJzY29wZXMiOlsibWVudS5pbnN1cmVkIiwibWVudS50YXNrYXBwcm92YWwiLCJtZW51LmxvZyIsIm1lbnUubWVtYmVyIiwibWVudS5maW5hbmNlIiwibWVudS5hcHBvaW50bWVudCIsIm1lbnUudXNlcm1hbmFnZW1lbnQiLCJtZW51Lm90aGVycyIsIm1lbnUudXRpbGlzYXRpb25yZXBvcnQiLCJtZW51LmJhdGNoIiwibWVudS5jb21tdW5pY2F0aW9uIiwibWVudS5jbGFpbXMiLCJtZW51LmRhc2NhbnBheSJdLCJyc2lkIjoiNTM1M2ZmMWUtOWYyMS00YTY0LThjNmQtMTU4NDM0YjNiYmM3IiwicnNuYW1lIjoiYWRtaW5tZW51In1dfSwic2NvcGUiOiJwcm9maWxlIGIyYi5hZG1pbnVzZXJzY29wZXMiLCJzaWQiOiJiYTVlMDIwNy02YjJkLTQ4ZjktOTk0My1mNDc3NTJiMTE3YWYiLCJpbnN1cmVybWFwIjpbIkdFIl0sInBob25lIjoiOTAxMTE3MjYiLCJuYW1lIjoiQWRtaW4gT3BzIEFkbWluIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiYWRtaW5vcHNAZ21haWwuY29tIiwic3RhcnRkYXRlIjoiMjAyMi0wMS0wMSIsImdpdmVuX25hbWUiOiJBZG1pbiBPcHMiLCJmYW1pbHlfbmFtZSI6IkFkbWluIiwiZW1haWwiOiJhZG1pbm9wc0BnbWFpbC5jb20ifQ.LI80oWrPtE4QVhQdtKK3v_stvnxNE43HFeK9vHaqt83uvtkzAg7WRRa7uBWn2q-o7rPSFg8Utzl5ccnlgq7jBy3WmmM-Kro3N3pP8FnvLrLv9X8i1QeE0PO7R5EWDw7waiIIi9CiOJzwJtLDtcGYo0hFdO-8pFKQY7dd_u5G1yyZIVp6bnRP8TI-fTHsref9HS5pkFzwxZsOPhAtUUqCfEg3FZ6FDyvbb6yXZ5OIIobM8xGFlCXFZ1yDK0FrsQSw8eYuNatc3K0H0MLUj5AzkATFhPt-zqeKDd20NoOU7d_mdNw4HT-6eKyUWDV7oeSH-iu2mrW28l7WrS1sZUSC6G"

	p, e := VerifyToken(tk)
	assert.NoError(t, e)
	fmt.Println(p)
}