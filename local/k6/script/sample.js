/**
 * ref. https://grafana.com/docs/k6/latest/using-k6
 */
import http from 'k6/http';
import {sleep} from 'k6';

const BASE_URL = "http://host.docker.internal:8080";
export const options = {
  vus: 200,
  duration: '60s',
};

export default function () {
  http.get(`${BASE_URL}/api`, {
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer jwt-token'
    },
  });
  sleep(0.1);
}