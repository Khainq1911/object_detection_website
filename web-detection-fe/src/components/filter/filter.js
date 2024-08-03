import React, { useState } from "react";
import classNames from "classnames/bind";
import styles from "./filter.modules.scss";
import Menu from "~/components/menu";
import "react-datepicker/dist/react-datepicker.css";
import DatePicker from "react-datepicker";
import * as apiServices from "~/services/homeService";
import { useLocation, useNavigate } from "react-router-dom";

const cx = classNames.bind(styles);

const dataFilter = {
  status: ["All", "Accept", "Reject"],
  cam: ["all", "0001", "0002", "0003"],
  type: ["All", "Human", "Vehicle"],
};

function Filter({ getFilterData, handleCloseFilter }) {
  const [openIndex, setOpenIndex] = useState(null);
  const [dateRange, setDateRange] = useState([null, null]);
  const [camera, setCamera] = useState("all");
  const [status, setStatus] = useState("All");
  const [event, setEvent] = useState("All");
  const navigate = useNavigate();

  const [timeFrom, timeTo] = dateRange;

  const toggleDropdown = (index) => {
    setOpenIndex(openIndex === index ? null : index);
  };

  const getCamValue = (val) => setCamera(val);
  const getStatusValue = (val) => setStatus(val);
  const getEventValue = (val) => setEvent(val);

  const filterData = async () => {
    try {
      if (timeFrom && !timeTo) {
        alert("Please choose an end date.");
        return;
      }

      const eventFilter =
        event === "All"
          ? "all"
          : event === "Human"
          ? "human_event"
          : "vehicle_event";

      const statusFilter =
        status === "All" ? "all" : status === "Accept" ? "true" : "false";

      const timeFromISO = timeFrom ? timeFrom.toISOString() : "all";
      const timeToISO = timeTo ? timeTo.toISOString() : "all";

        
      const res = await apiServices.filterData(
        eventFilter,
        timeFromISO,
        timeToISO,
        camera,
        statusFilter
      );

      
      const newPath = `/filter/${eventFilter}/${timeFromISO}/${timeToISO}/${camera}/${statusFilter}`;
      
    
      navigate(newPath, { replace: true });
      
     
      getFilterData(res.data);
      handleCloseFilter();
    } catch (error) {
      console.error("Failed to fetch filter data:", error);
      alert("An error occurred while fetching filter data.");
    }
  };

  return (
    <div className={cx("filter_container")}>
      <div className={cx("cam_menu")}>
        <p>Camera:</p>
        <Menu
          data={dataFilter.cam}
          isOpen={openIndex === 0}
          onToggle={() => toggleDropdown(0)}
          getValue={getCamValue}
        />
      </div>
      <div className={cx("type_menu")}>
        <p>Event type:</p>
        <Menu
          data={dataFilter.type}
          isOpen={openIndex === 1}
          onToggle={() => toggleDropdown(1)}
          getValue={getEventValue}
        />
      </div>
      <div className={cx("status_menu")}>
        <p>Status:</p>
        <Menu
          data={dataFilter.status}
          isOpen={openIndex === 2}
          onToggle={() => toggleDropdown(2)}
          getValue={getStatusValue}
        />
      </div>
      <div className={cx("calendar")}>
        <p>Time:</p>
        <DatePicker
          selected={timeFrom}
          onChange={(update) => setDateRange(update)}
          startDate={timeFrom}
          endDate={timeTo}
          selectsRange
        />
      </div>
      <div>
        <button className={cx("filter_menu")} onClick={filterData}>
          Filter
        </button>
      </div>
    </div>
  );
}

export default Filter;
