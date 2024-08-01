import HomeCard from "~/components/HomeCard";
import { useEffect, useState } from "react";
import * as apiServices from "~/services/homeService";
import classNames from "classnames/bind";
import styles from "./home.module.scss";
import Modal from "~/components/modal";
import ModalConfirm from "~/components/modalConfirm";
import Filter from "~/components/filter";

const cx = classNames.bind(styles);
function Home() {
  const [data, setData] = useState([]);
  const [modalMessage, setModalMessage] = useState(null);
  const [activeModal, setActiveModal] = useState(false);
  const [activeConfirm, setActiveConfirm] = useState(false);
  const [confirmMessage, setConfirmMessage] = useState(null);
  const [action, setAction] = useState();

  const fetchData = async () => {
    try {
      const res = await apiServices.getMessage();
      setData(res.data);
      return res.data;
    } catch (error) {
      console.log(error);
    }
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
      if (activeModal) {
        setActiveModal(!activeModal);
      }
    } catch (error) {
      console.error(error);
    }
  };
  useEffect(() => {
    fetchData();
    const shortPolling = setInterval(() => {
      fetchData();
    }, 50000);
    return () => {
      clearInterval(shortPolling);
    };
  }, []);

  const handleUpdateData = async () => {
    await fetchData();
  };

  const handleGetMessage = (message) => {
    setModalMessage(message);
    setConfirmMessage(message);
  };
  const handleActiveModal = () => {
    setActiveModal(!activeModal);
    document.body.style.overflow = "hidden";
  };
  const handleCloseModal = () => {
    setActiveModal(!activeModal);
    document.body.style.overflow = "auto";
  };

  const handleActiveConfirm = () => {
    setActiveConfirm(!activeConfirm);
  };
  const handleCloseConfirm = () => {
    setActiveConfirm(!activeConfirm);
  };
  const handleSetAction = (action) => {
    setAction(action);
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
              handleActiveConfirm={handleActiveConfirm}
              handleSetAction={handleSetAction}
            />
          </div>
        );
      })}
      {activeModal && (
        <Modal
          data={modalMessage}
          handleCloseModal={handleCloseModal}
          handleMessageAction={handleMessageAction}
          handleActiveConfirm={handleActiveConfirm}
          handleSetAction={handleSetAction}
        />
      )}
      {activeConfirm && <div className={cx("confirm")}></div>}
      {activeConfirm && (
        <ModalConfirm
          data={confirmMessage}
          action={action}
          handleCloseConfirm={handleCloseConfirm}
          handleCloseModal={handleActiveModal}
          handleMessageAction={handleMessageAction}
        />
      )}
      <Filter />
    </div>
  );
}
export default Home;
