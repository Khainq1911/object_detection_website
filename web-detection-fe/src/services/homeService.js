import api from "./api";

export const getMessage = async () => {
  const response = await api.get("/");
  return response.data;
};
export const acceptMessage = async (messageId) => {
  const response = await api.get(`/accept/${messageId}`);
  return response;
};
export const rejectMessage = async (messageId) => {
  const response = await api.get(`/reject/${messageId}`);
  return response;
};
export const discardAckMessage = async (messageId) => {
  const response = await api.get(`/discard-ack/${messageId}`);
  return response;
};
export const filterData = async (
  eventType,
  timeFrom,
  timeTo,
  cameraID,
  status
) => {
  const response = await api.get(
    `filter/${eventType}/${timeFrom}/${timeTo}/${cameraID}/${status}`
  );
  return response.data;
};
