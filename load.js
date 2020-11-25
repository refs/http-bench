import http from 'k6/http';
import { check } from 'k6';

export let options = {
    vus: 100,
    iterations: 1000000,
    insecureSkipTLSVerify: true,
};

export default function () {
    let base = 'http://localhost:8082'

    let res = http.get(base)

    check(res, {
            'response code was 200': (res) => res.status == 200
        }
    )
}
