import React, { useState } from 'react';
import classNames from 'classnames/bind';
import styles from './filter.modules.scss';
import Menu from '~/components/menu';

const cx = classNames.bind(styles);

const dataFilter = {
  status: ['All', 'Accept', 'Reject'],
  cam: ['All', '0001', '0002', '0003'],
  type: ['All', 'Human', 'Vehicle'],
};

function Filter() {
  const [openIndex, setOpenIndex] = useState(null);

  const toggleDropdown = (index) => {
    setOpenIndex(openIndex === index ? null : index);
  };

  return (
    <div className={cx('filter_container')}>
      <div className={cx('cam_menu')}>
        <p>Camera:</p>
        <Menu
          
          data={dataFilter.cam}
          isOpen={openIndex === 0}
          onToggle={() => toggleDropdown(0)}
        />
      </div>
      <div className={cx('type_menu')}>
        <p>Event type:</p>
        <Menu
         
          data={dataFilter.type}
          isOpen={openIndex === 1}
          onToggle={() => toggleDropdown(1)}
        />
      </div>
      <div className={cx('status_menu')}>
        <p>Status:</p>
        <Menu
         
          data={dataFilter.status}
          isOpen={openIndex === 2}
          onToggle={() => toggleDropdown(2)}
        />
      </div>
      <div>
        <button className={cx('filter_menu')}>Filter</button>
      </div>
    </div>
  );
}

export default Filter;
