import axios from 'axios';

const host = 'http://127.0.0.1:8080/api';

const request = axios.create({
  baseURL: host,
  timeout: 1000,
  withCredentials: true,
});

request.interceptors.request.use((config) => config);

request.defaults.timeout = 5000;

request.interceptors.response.use(
  (response) => {
    if (response.status === 200) {
      const { data } = response;
      console.log('axios:', data);

      if (data.code === 200) {
        return data;
      }
    }
    return response;
  },
  (err) => {
    const { config } = err;

    if (!config || !config.retry) return Promise.reject(err);

    config.retryCount = config.retryCount || 0;

    if (config.retryCount >= config.retry) {
      return Promise.reject(err);
    }

    config.retryCount += 1;

    const backoff = new Promise((resolve) => {
      setTimeout(() => {
        resolve(null);
      }, config.retryDelay || 1);
    });

    return backoff.then(() => request(config));
  }
);

export default request;
