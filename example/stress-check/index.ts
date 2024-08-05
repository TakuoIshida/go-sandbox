import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  vus: 30,
  duration: "10s",
};

export default function () {
  const url = "http://localhost:8080";
  const path = "/api/v1/company/12/scope1n2/upload-files";
  const X_AMZN_OIDC_DATA = "";
  const X_AMZN_OIDC_ID = "";
  const X_AMZN_OIDC_TOKEN = "";
  const params = {
    headers: {
      "Content-Type": "application/json",
      "x-amzn-oidc-data": X_AMZN_OIDC_DATA,
      "x-amzn-oidc-identity": X_AMZN_OIDC_ID,
      "x-amzn-oidc-accesstoken": X_AMZN_OIDC_TOKEN,
    },
  };
  let res = http.get(url + path, params);
  check(res, { "success login": (r) => r.status === 200 });
}
