import unitApi from "../api/api";

const unitSrv = {
  list: async () => {
    const { data } = await unitApi.list();
    return data;
  },
};

export default unitSrv;
