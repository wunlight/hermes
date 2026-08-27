import httpClient from "@/core/utils/http-client";

const productApi = {
  list: () => httpClient.get("product"),
};

export default productApi;
