import axiosClient from "./axiosClient";
const authApi = {
  signup() {
    const url = "/generateMnemonicPhrase";
    return axiosClient.get(url);
  },

  login(wordList) {
    const url = "/validateMnemonicPhrase";
    return axiosClient.post(url, wordList);
  },
};

export default authApi;
