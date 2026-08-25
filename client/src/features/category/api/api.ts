import httpClient from "@/core/utils/http-client";

const categoryApi = {
  list: () => httpClient.get("category"),
};

export default categoryApi;
