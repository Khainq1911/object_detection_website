import classNames from "classnames/bind";
import styles from "./modalConfirm.module.scss";

const cx = classNames.bind(styles);

function ModalConfirm({
  action,
  data,
  handleCloseConfirm,
  handleCloseModal,
  handleMessageAction,
}) {
  const messageId = data.message_id;
  return (
    <div className={cx("modalConfirm")}>
      <p className={cx("text-modal")}>Do you want to {`${action}`}?</p>
      <div className={cx("confirm_btn")}>
        <button
          onClick={() => {
            handleCloseConfirm();
            handleMessageAction(action, messageId);
          }}
        >
          Yes
        </button>
        <button
          onClick={() => {
            handleCloseConfirm();
          }}
        >
          No
        </button>
      </div>
    </div>
  );
}

export default ModalConfirm;
