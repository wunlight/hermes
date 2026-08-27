import brandApi from "../api/api";

const brandSrv = {
  list: async () => {
    const { data } = await brandApi.list();
    return data;
  },
};

export default brandSrv;
