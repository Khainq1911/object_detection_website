import classNames from "classnames/bind";
import styles from "./chart.module.scss";
import PieChart from "~/components/chart/pieChart";
import { useEffect, useState } from "react";
import * as apiServices from "~/services/homeService";

const cx = classNames.bind(styles);

function Chart() {
  const [status, setStatus] = useState({
    labels: ["Accepted", "Rejected", "Pending"],
    datasets: [
      {
        label: "Status Chart",
        data: [0, 0, 0],
        backgroundColor: ["#36A2EB", "#FF6384", "#FFCE56"],
        hoverBackgroundColor: ["#36A2EB", "#FF6384", "#FFCE56"],
      },
    ],
  });

  const [object, setObject] = useState({
    labels: ["Human", "Vehicle"],
    datasets: [
      {
        label: "Object Chart",
        data: [0, 0],
        backgroundColor: ["#36A2EB", "#FF6384"],
        hoverBackgroundColor: ["#36A2EB", "#FF6384"],
      },
    ],
  });

  const [camera, setCamera] = useState({
    labels: ["0001", "0002", "0003"],
    datasets: [
      {
        label: "Camera Chart",
        data: [0, 0, 0],
        backgroundColor: ["#36A2EB", "#FF6384", "#FFCE56"],
        hoverBackgroundColor: ["#36A2EB", "#FF6384", "#FFCE56"],
      },
    ],
  });

  const countStatus = (messages) => {
    let accept = 0,
      reject = 0,
      pending = 0;

    messages.forEach((message) => {
      if (message.status === null) {
        pending++;
      } else if (message.status === true) {
        accept++;
      } else {
        reject++;
      }
    });

    setStatus((prevData) => ({
      ...prevData,
      datasets: [
        {
          ...prevData.datasets[0],
          data: [accept, reject, pending],
        },
      ],
    }));
  };

  const countCamera = (messages) => {
    let cam0001 = 0,
      cam0002 = 0,
      cam0003 = 0;

    messages.forEach((message) => {
      if (message.camera_id === "0001") {
        cam0001++;
      } else if (message.camera_id === "0002") {
        cam0002++;
      } else {
        cam0003++;
      }
    });

    setCamera((prev) => ({
      ...prev,
      datasets: [
        {
          ...prev.datasets[0],
          data: [cam0001, cam0002, cam0003],
        },
      ],
    }));
  };

  const countObject = (messages) => {
    let human = 0,
      vehicle = 0;
    const objects = messages.flatMap((message) => message.object_list);

    objects.forEach((object) => {
      if (object.object_type === "Human") {
        human++;
      } else {
        vehicle++;
      }
    });

    setObject((prev) => ({
      ...prev,
      datasets: [
        {
          ...prev.datasets[0],
          data: [human, vehicle],
        },
      ],
    }));
  };

  useEffect(() => {
    const getMessage = async () => {
      try {
        const res = await apiServices.getMessage();
        countStatus(res.data);
        countCamera(res.data);
        countObject(res.data);
      } catch (error) {
        console.log(error);
      }
    };
    getMessage();
    const updateData = setInterval(() => {
      getMessage();
    }, 60000);

    return () => clearInterval(updateData);
  }, []);

  return (
    <div className={cx("chart_container")}>
      <div className={cx("chart")}>
        <div className={cx("chart_title")}>Status Chart</div>
        <PieChart chartData={status} />
      </div>
      <div className={cx("chart")}>
        <div className={cx("chart_title")}>Object Chart</div>
        <PieChart chartData={object} />
      </div>
      <div className={cx("chart")}>
        <div className={cx("chart_title")}>Camera Chart</div>
        <PieChart chartData={camera} />
      </div>
    </div>
  );
}

export default Chart;
