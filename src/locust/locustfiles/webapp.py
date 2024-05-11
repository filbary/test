from locust import HttpUser, task, between

class WebappUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def convert(self):
        self.client.post("/login", json={"fahrenheit":42})
