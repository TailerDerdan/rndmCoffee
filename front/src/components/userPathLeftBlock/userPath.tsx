import React, { useEffect, useState } from "react";
import styles from "./userPath.module.css";
import { Link, useLocation } from "react-router-dom";
import { LineSeparator } from "../icons/icons";

export const UserPath = () => {
	const [classesForFirstElementList] = useState(
		styles.list__elementOfList + " " + styles.elementOfList__active,
	);
	const [classesForSecondElementList, setClassesForSecondElementList] =
		useState(styles.list__elementOfList);
	const [classesForThirdElementList, setClassesForThirdElementList] =
		useState(styles.list__elementOfList);

	const location = useLocation();
	const [keyLocation, setKeyLocation] = useState(location.pathname);
	useEffect(() => {
		setKeyLocation(location.pathname);
	}, [location.pathname]);

	if (keyLocation === "http://localhost:3000/auth_profile/contacts") {
		setClassesForSecondElementList(styles.list__elementOfList);
		setClassesForThirdElementList(styles.list__elementOfList);
	}

	if (keyLocation === "http://localhost:3000/auth_profile/activity") {
		setClassesForSecondElementList(
			styles.list__elementOfList + " " + styles.elementOfList__active,
		);
		setClassesForThirdElementList(styles.list__elementOfList);
	}

	if (keyLocation === "http://localhost:3000/auth_profile/welcome") {
		setClassesForSecondElementList(
			styles.list__elementOfList + " " + styles.elementOfList__active,
		);
		setClassesForThirdElementList(
			styles.list__elementOfList + " " + styles.elementOfList__active,
		);
	}

	return (
		<div className={styles.wrapper__userPath}>
			<div className={styles.userPath__list}>
				<ol className={styles.list}>
					<Link
						to={"/auth_profile/contacts"}
						className={classesForFirstElementList}
						onClick={() => {
							setClassesForSecondElementList(
								styles.list__elementOfList,
							);
							setClassesForThirdElementList(
								styles.list__elementOfList,
							);
						}}
					>
						Персональная информация
					</Link>
					<div className={styles.lineSeparator}>
						{<LineSeparator />}
					</div>
					<Link
						to={"/auth_profile/activity"}
						className={classesForSecondElementList}
						onClick={() => {
							setClassesForSecondElementList(
								styles.list__elementOfList +
									" " +
									styles.elementOfList__active,
							);
							setClassesForThirdElementList(
								styles.list__elementOfList,
							);
						}}
					>
						Ваши интересы
					</Link>
					<div className={styles.lineSeparator}>
						{<LineSeparator />}
					</div>
					<Link
						to={"/auth_profile/welcome"}
						className={classesForThirdElementList}
						onClick={() => {
							setClassesForSecondElementList(
								styles.list__elementOfList +
									" " +
									styles.elementOfList__active,
							);
							setClassesForThirdElementList(
								styles.list__elementOfList +
									" " +
									styles.elementOfList__active,
							);
						}}
					>
						Добро пожаловать!
					</Link>
				</ol>
			</div>
		</div>
	);
};
