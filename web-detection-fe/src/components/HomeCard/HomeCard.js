import classNames from "classnames/bind";
import styles from "./homeCard.module.scss";

const cx = classNames.bind(styles);

function HomeCard({
  data,
  className,
  handleGetMessage,
  handleActiveModal,
  handleActiveConfirm,
  handleSetAction,
}) {
  return (
    <div className={cx("homeCard_container", className)}>
      <img src={data.image_url} className={cx("card_image")} alt="Card" />
      <div className={cx("card_text")}>
        <p className={cx("message")}>{data.message_id}</p>
        <div className={cx("time", "common")}>
          <p>Timestamp:</p>
          <p>{data.timestamp}</p>
        </div>
        <div className={cx("location", "common")}>
          <p>Location ID:</p>
          <p>{data.location_id}</p>
        </div>
        <div className={cx("model", "common")}>
          <p>Model:</p>
          <p>{data.description_mdl}</p>
        </div>
        <div className={cx("camera", "common")}>
          <p>Camera:</p>
          <p>{data.description_cam}</p>
        </div>
        <div className={cx("camera_id", "common")}>
          <p>Camera ID:</p>
          <p>{data.camera_id}</p>
        </div>
        <div className={cx("camera_type", "common")}>
          <p>Camera Type:</p>
          <p>{data.type}</p>
        </div>
        <div className={cx("object", "common")}>
          <p>Number of Objects:</p>
          <p>{data.number_of_objects}</p>
        </div>
        <div className={cx("event", "common")}>
          <p>Number of Events:</p>
          <p>{data.number_of_event}</p>
        </div>
        <div className={cx("status", "common")}>
          <p>Status:</p>
          <p>
            {data.status === null
              ? "Pending"
              : data.status === true
              ? "Accepted"
              : "Rejected"}
          </p>
        </div>
      </div>
      <div className={cx("card_btn")}>
        <button
          className={cx("watch_btn")}
          onClick={() => {
            handleActiveModal();
            handleGetMessage(data);
          }}
        >
          Watch Video
        </button>
        {data.status === null ? (
          <div>
            <button
              className={cx("accept_btn")}
              onClick={() => {
                handleGetMessage(data);
                handleSetAction("accept");
                handleActiveConfirm();
              }}
            >
              Accept
            </button>
            <button
              className={cx("reject_btn")}
              onClick={() => {
                handleGetMessage(data);
                handleSetAction("reject");
                handleActiveConfirm();
              }}
            >
              Reject
            </button>
          </div>
        ) : (
          <button
            className={cx("discard_btn")}
            onClick={() => {
              handleGetMessage(data);
              handleSetAction("discardAck");
              handleActiveConfirm();
            }}
          >
            Discard-Ack
          </button>
        )}
      </div>
    </div>
  );
}

export default HomeCard;
