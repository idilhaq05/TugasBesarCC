import http from 'k6/http';
import { sleep, check } from 'k6';

export let options = {
  stages: [
    { duration: '1m', target: 50 },
    { duration: '1m', target: 200 },
    { duration: '1m', target: 500 },
    { duration: '1m', target: 1000 },
    { duration: '1m', target: 10000 },
    { duration: '1m', target: 50000 },
    { duration: '1m', target: 100000 },
    { duration: '2m', target: 100000 },
    { duration: '2m', target: 0 },
  ],
};

const BASE_URL = 'http://192.168.231.128';

export function setup() {
  let res = http.post(`${BASE_URL}/api/login`, {
    email: 'admin@ecosteps.com',
    password: 'admin',
  });

  check(res, { 'login ok': (r) => r.status === 200 });

  return { token: res.json('token') };
}

export default function (data) {
  let params = {
    headers: { Authorization: `Bearer ${data.token}` },
  };

  let res = http.get(`${BASE_URL}/api/dashboard`, params);

  check(res, {
    'dashboard ok': (r) => r.status === 200,
    'fast < 1s': (r) => r.timings.duration < 1000,
  });

  sleep(2);
}
