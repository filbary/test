"""Module to perform load testing with Locust."""

from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    """A user class for quick start load testing scenarios."""
    wait_time = between(1, 2)

    @task
    def hello_world(self):
        """Accesses the convert endpoint to simulate a user interaction."""
        self.client.get("/convert/43")
