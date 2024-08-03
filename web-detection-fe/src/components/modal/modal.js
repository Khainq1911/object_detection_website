import classNames from "classnames/bind";
import styles from "./modal.module.scss";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faX } from "@fortawesome/free-solid-svg-icons";

const cx = classNames.bind(styles);
function Modal({
  data,
  handleSetAction,
  handleCloseModal,
  handleActiveConfirm,
  handleMessageAction,
}) {
  return (
    <div className={cx("modal_container")}>
      <div className={cx("modal_header")}>
        <p>{data.message_id}</p>
        <FontAwesomeIcon
          icon={faX}
          className={cx("icon")}
          onClick={() => {
            handleCloseModal();
          }}
        />
      </div>
      <video
        width="1400"
        height="500"
        controls
        autoPlay
        className={cx("video")}
      >
        <source src={data.video_url} type="video/mp4" />
        Your browser does not support the video tag.
      </video>
      <div className={cx("header_infor")}>
        {data.object_list.map((object, index) => {
          return (
            <div className={cx("infor_wrapper")}>
              <div className={cx("common")}>
                <p>Object id:</p>
                <p>{object.object_id}</p>
              </div>
              <div className={cx("common")}>
                <p>Object type:</p>
                <p>{object.object_type}</p>
              </div>
              {object.object_type === "Human" ? (
                <div>
                  <div className={cx("common")}>
                    <p>Gender:</p>
                    <p>{object.gender}</p>
                  </div>
                  <div className={cx("common")}>
                    <p>Age:</p>
                    <p>{object.age}</p>
                  </div>
                </div>
              ) : (
                <div>
                  <div className={cx("common")}>
                    <p>Vehicle type:</p>
                    <p>{object.vehicle_type}</p>
                  </div>
                  <div className={cx("common")}>
                    <p>Vehicle brand:</p>
                    <p>{object.vehicle_brand}</p>
                  </div>
                  <div className={cx("common")}>
                    <p>Vehicle color:</p>
                    <p>{object.vehicle_color}</p>
                  </div>
                  <div className={cx("common")}>
                    <p>Vehicle license:</p>
                    <p>{object.vehicle_license}</p>
                  </div>
                </div>
              )}
              <div className={cx("common")}>
                <p>bbox TopLeftX:</p>
                <p>{object.bbox_topleftx}</p>
              </div>
              <div className={cx("common")}>
                <p>bbox TopLeftY:</p>
                <p>{object.bbox_toplefty}</p>
              </div>
              <div className={cx("common")}>
                <p>bbox BottomRightX:</p>
                <p>{object.bbox_bottomrightx}</p>
              </div>
              <div className={cx("common")}>
                <p>bbox BottomRightY:</p>
                <p>{object.bbox_bottomrighty}</p>
              </div>
            </div>
          );
        })}
      </div>
      <div className={cx("modal_btn_wrapper")}>
        <button
          className={cx("accept_btn")}
          onClick={() => {
            handleSetAction("accept");

            handleActiveConfirm();
          }}
        >
          Accept
        </button>
        <button
          className={cx("reject_btn")}
          onClick={() => {
            handleSetAction("reject");

            handleActiveConfirm();
          }}
        >
          Reject
        </button>
      </div>
    </div>
  );
}
export default Modal;
