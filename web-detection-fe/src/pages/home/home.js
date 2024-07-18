import HomeCard from "~/components/HomeCard";
import { useEffect, useState } from "react";
import * as apiServices from "~/services/homeService";
import classNames from "classnames/bind";
import styles from "./home.module.scss";
import Modal from "~/components/modal";

const cx = classNames.bind(styles);
function Home() {
  const [data, setData] = useState([]);
  const [modalMessage, setModalMessage] = useState(null);
  const [activeModal, setActiveModal] = useState(false);

  const fetchData = async () => {
    try {
      const res = await apiServices.getMessage();
      setData(res.data);
      return res.data;
    } catch (error) {
      console.log(error);
    }
  };
  const handleUpdateData = async () => {
    await fetchData();
  };

  const handleMessageAction = async (action, messageId) => {
    try {
      let res;
      if (action === "accept") {
        res = await apiServices.acceptMessage(messageId);
      } else if (action === "reject") {
        res = await apiServices.rejectMessage(messageId);
      } else if (action === "discardAck") {
        res = await apiServices.discardAckMessage(messageId);
      }
      if (res) {
        handleUpdateData();
      }
    } catch (error) {
      console.error(error);
    }
  };
  useEffect(() => {
    fetchData();
    const shortPolling = setInterval(() => {
      fetchData();
    }, 5000);
    return () => {
      clearInterval(shortPolling);
    };
  }, []);

  const handleGetMessage = (message) => {
    setModalMessage(message);
  };
  const handleActiveModal = () => {
    setActiveModal(!activeModal);
    document.body.style.overflow = "hidden";
  };
  const handleCloseModal = () => {
    setActiveModal(!activeModal);
    document.body.style.overflow = "auto";
  };
  return (
    <div className={cx("home_container")}>
      {activeModal && <div className={cx("overlay")}></div>}
      {data.map((message, index) => {
        return (
          <div key={index} className={cx("card_wrap")}>
            <HomeCard
              data={message}
              className={cx("home_card")}
              handleGetMessage={() => handleGetMessage(message)}
              handleActiveModal={handleActiveModal}
              handleUpdateData={handleUpdateData}
              handleMessageAction={handleMessageAction}
            />
          </div>
        );
      })}
      {activeModal && (
        <Modal
          data={modalMessage}
          handleCloseModal={handleCloseModal}
          handleMessageAction={handleMessageAction}
        />
      )}
    </div>
  );
}
export default Home;
