import axiosClient from "./axiosClient";
const transactionApi = {
  getAddressData() {
    const url = `/address/${localStorage.getItem("publicKey")}`;
    return axiosClient.get(url);
  },

  createPendingTransaction({ amount, sender, recipient }) {
    const url = "/transaction/broadcast";
    return axiosClient.post(url, { amount, sender, recipient });
  },
};

export default transactionApi;
