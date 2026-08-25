import categoryApi from "../api/api";

const categorySrv = {
  list: async () => {
    const { data } = await categoryApi.list();
    return data;
  },
};

export default categorySrv;
