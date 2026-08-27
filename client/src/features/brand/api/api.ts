import httpClient from "@/core/utils/http-client";

const brandApi = {
  list: () => httpClient.get("brand"),
};

export default brandApi;
