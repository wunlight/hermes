import productApi from "../api/api";

const productSrv = {
  list: async () => {
    const { data } = await productApi.list();
    return data;
  },
};

export default productSrv;
