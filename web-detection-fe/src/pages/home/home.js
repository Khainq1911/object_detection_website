import HomeCard from "~/components/HomeCard";
import { useEffect, useState } from "react";
import * as apiServices from "~/services/homeService";
import classNames from "classnames/bind";
import styles from "./home.module.scss";
import Modal from "~/components/modal";

const cx = classNames.bind(styles);
function Home() {
  const [data, setData] = useState([]);
  const [modalMessage, setModalMessage] = useState([]);
  const [activeModal, setActiveModel] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await apiServices.getMessage();
        setData(res.data);
        return res.data;
      } catch (error) {
        console.log(error);
      }
    };
    fetchData();
  }, []);

  const handleGetMessage = (message) => {
    setActiveModel(!activeModal);
    setModalMessage(message);
    document.body.style.overflow = "hidden";
  };

  return (
    <div className={cx("home_container")}>
      {activeModal && <div className={cx("overlay")}></div>}
      {data.map((message, index) => {
        return (
          <div
            key={index}
            className={cx("card_wrap")}
            onClick={() => handleGetMessage(message)}
          >
            <HomeCard key={index} data={message} className={cx("home_card")} />
          </div>
        );
      })}
      {activeModal && <Modal data={modalMessage} />}
    </div>
  );
}
export default Home;
