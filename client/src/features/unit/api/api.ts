import httpClient from "@/core/utils/http-client";

const unitApi = {
  list: () => httpClient.get("unit"),
};

export default unitApi;
