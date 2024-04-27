import React from "react";
import { Outlet } from "react-router-dom";
import styles from "./mainPage.module.css";
import { ButtonTextLink } from "../../components/button/button";

export const MainPage = () => {
	return (
		<>
			<header className={styles.mainPage__header}>
				<div className={styles.header__buttonHome}>
					<ButtonTextLink
						id={"home"}
						link={"/main/home"}
						title={"Главная"}
					/>
				</div>
				<div className={styles.header__buttonProfile}>
					<ButtonTextLink
						id={"profile"}
						link={"/main/profile"}
						title={"Профиль"}
					/>
				</div>
			</header>
			<div className={styles.mainContent}>
				<Outlet />
			</div>
		</>
	);
};
