import React from "react";
import styles from "./mainProfilePage.module.css";
import { Outlet } from "react-router-dom";
import { UserPath } from "../../../components/userPathLeftBlock/userPath";

export const MainProfile = () => {
	return (
		<div className={styles.background}>
			<div className={styles.wrapper__auth}>
				<div className={styles.wrapper__userPath}>{<UserPath />}</div>
				<Outlet />
			</div>
		</div>
	);
};
