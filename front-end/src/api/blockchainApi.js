import axiosClient from "./axiosClient";
const blockchainApi = {
  getBlockChain() {
    const url = "/blockchain";
    return axiosClient.get(url);
  },
};

export default blockchainApi;
