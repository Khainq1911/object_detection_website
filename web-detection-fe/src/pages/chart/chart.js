import Modal from "~/components/modal";
import classNames from "classnames/bind";
import styles from "./chart.module.scss";

const cx = classNames.bind(styles);

function Chart() {
  return (
    <div className={cx("chart_container")}>
      <Modal />
    </div>
  );
}
export default Chart;
