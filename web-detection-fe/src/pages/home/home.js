import HomeCard from "~/components/HomeCard";
import { useEffect, useState } from "react";
import * as apiServices from "~/services/homeService";
import classNames from "classnames/bind";
import styles from "./home.module.scss";

const cx = classNames.bind(styles);
function Home() {
  const [data, setData] = useState([]);
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
  return (
    <div className={cx("home_container")}>
      {data.map((message, index) => {
        return (
          <div className={cx("card_wrap")}>
            <HomeCard key={index} data={message} className={cx("home_card")} />
          </div>
        );
      })}
    </div>
  );
}
export default Home;
