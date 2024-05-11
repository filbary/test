from locust import HttpUser, task, between

class ServingUser(HttpUser):
    wait_time = between(1, 2)

    @task
    def convert(self):
        self.client.post("/v1/models/model:predict", json={"inputs": 42})
