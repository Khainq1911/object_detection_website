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
  const [activeModal, setActiveModel] = useState(false);

  const fetchData = async () => {
    try {
      const res = await apiServices.getMessage();
      setData(res.data);
      return res.data;
    } catch (error) {
      console.log(error);
    }
  };
  const handleUpdateData = () => {
    fetchData();
  };
  useEffect(() => {
    /* const shortPolling = setInterval(() => {
      fetchData();
    }, 5000);
    return () => {
      clearInterval(shortPolling);
    }; */
    fetchData();
  }, [data]);

  const handleGetMessage = (message) => {
    setModalMessage(message);
  };
  const handleActiveModal = () => {
    setActiveModel(!activeModal);
    document.body.style.overflow = "hidden";
  };
  const handleCloseModal = () => {
    setActiveModel(!activeModal);
    document.body.style.overflow = "auto";
  };
  return (
    <div className={cx("home_container")}>
      {activeModal && <div className={cx("overlay")}></div>}
      {data.map((message, index) => {
        return (
          <div key={index} className={cx("card_wrap")}>
            <HomeCard
              key={index}
              data={message}
              className={cx("home_card")}
              handleGetMessage={() => handleGetMessage(message)}
              handleActiveModal={handleActiveModal}
              handleUpdateData={handleUpdateData}
            />
          </div>
        );
      })}
      {activeModal && (
        <Modal data={modalMessage} handleCloseModal={handleCloseModal} />
      )}
    </div>
  );
}
export default Home;
