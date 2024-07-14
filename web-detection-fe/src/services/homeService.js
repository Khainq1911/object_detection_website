import api from "./api";

export const getMessage = async () => {
  const response = await api.get("/");
  return response.data;
};
