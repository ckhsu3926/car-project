-- MySQL Script generated by MySQL Workbench
-- Mon Jul 25 15:04:08 2022
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema car-project
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `car-project` ;

-- -----------------------------------------------------
-- Schema car-project
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `car-project` DEFAULT CHARACTER SET utf8mb4 ;
USE `car-project` ;

-- -----------------------------------------------------
-- Table `car-project`.`user`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `car-project`.`user` ;

CREATE TABLE IF NOT EXISTS `car-project`.`user` (
  `id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '車主ID',
  `username` VARCHAR(16) NOT NULL COMMENT '車主帳號',
  `password` VARCHAR(45) NOT NULL COMMENT '登入密碼',
  `real_name` VARCHAR(45) NULL DEFAULT '' COMMENT '使用者姓名',
  `address` VARCHAR(45) NULL DEFAULT '' COMMENT '地址',
  `login_token` VARCHAR(45) NULL DEFAULT '' COMMENT '登入權杖',
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = '使用者(車主)';


-- -----------------------------------------------------
-- Table `car-project`.`vehicle`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `car-project`.`vehicle` ;

CREATE TABLE IF NOT EXISTS `car-project`.`vehicle` (
  `id` SMALLINT UNSIGNED NOT NULL COMMENT '車輛ID',
  `user_id` SMALLINT UNSIGNED NOT NULL COMMENT '車主ID',
  `name` VARCHAR(20) NOT NULL COMMENT '車輛名稱',
  `license` VARCHAR(20) NOT NULL COMMENT '車牌號碼',
  `compony` VARCHAR(45) NULL DEFAULT '' COMMENT '廠牌',
  `model` VARCHAR(45) NULL DEFAULT '' COMMENT '款式',
  `engine_displacement` INT(5) UNSIGNED NULL DEFAULT 0 COMMENT '排氣量',
  `engine_number` VARCHAR(45) NULL DEFAULT '' COMMENT '引擎號碼',
  `default_octane_number` INT(4) UNSIGNED NULL DEFAULT 95 COMMENT '預設辛烷值',
  `purchase` INT(10) UNSIGNED ZEROFILL NULL DEFAULT 0 COMMENT '購入金額',
  `purchase_date` VARCHAR(10) NULL DEFAULT '' COMMENT '購入日期, format: 2021/12/31',
  `purchase_location` VARCHAR(45) NULL DEFAULT '' COMMENT '購入地點',
  `purchase_mileage` DECIMAL(7,2) UNSIGNED NULL DEFAULT 0.0 COMMENT '購入里程',
  `sold` INT(10) UNSIGNED NULL COMMENT '售出金額',
  `sold_date` VARCHAR(10) NULL COMMENT '售出日期, format: 2021/12/31',
  `sold_mileage` DECIMAL(7,2) UNSIGNED NULL COMMENT '售出里程',
  `mileage_reset` DECIMAL(7,2) UNSIGNED NULL COMMENT '儀表更換前最後里程',
  PRIMARY KEY (`id`),
  INDEX `INDEX_User_Vehicle` USING BTREE (`id`, `user_id`))
ENGINE = InnoDB
COMMENT = '車輛';


-- -----------------------------------------------------
-- Table `car-project`.`vehicle_insurance`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `car-project`.`vehicle_insurance` ;

CREATE TABLE IF NOT EXISTS `car-project`.`vehicle_insurance` (
  `id` VARCHAR(20) NOT NULL COMMENT '保單號碼',
  `vehicle_id` SMALLINT UNSIGNED NOT NULL COMMENT '所屬車輛ID',
  `name` VARCHAR(45) NOT NULL COMMENT '保單名稱',
  `compony` VARCHAR(45) NOT NULL COMMENT '保險公司',
  `start_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '保險生效日期',
  `end_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '結束日期',
  `value` INT(7) UNSIGNED NOT NULL COMMENT '保單金額',
  `insured` VARCHAR(45) NOT NULL COMMENT '被保險人',
  `guarantor` VARCHAR(45) NOT NULL COMMENT '要保人',
  PRIMARY KEY (`id`),
  INDEX `INDEX_Vehicle` USING BTREE (`vehicle_id`))
ENGINE = InnoDB
COMMENT = '車輛保險';


-- -----------------------------------------------------
-- Table `car-project`.`insurance`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `car-project`.`insurance` ;

CREATE TABLE IF NOT EXISTS `car-project`.`insurance` (
  `id` TINYINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `vehicle_insurance_id` VARCHAR(20) NOT NULL COMMENT '保單號碼',
  `type` VARCHAR(45) NULL DEFAULT '' COMMENT '保險種類',
  `condition` VARCHAR(45) NULL DEFAULT '' COMMENT '保險條件(每)',
  `upper_limit` INT(10) UNSIGNED NULL DEFAULT 0 COMMENT '賠償上限',
  `deductible` INT(10) UNSIGNED ZEROFILL NULL DEFAULT 0 COMMENT '自負額',
  `value` INT(10) UNSIGNED ZEROFILL NOT NULL COMMENT '保險費',
  PRIMARY KEY (`id`, `vehicle_insurance_id`))
ENGINE = InnoDB
COMMENT = '保單內容';


-- -----------------------------------------------------
-- Table `car-project`.`refueling_record`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `car-project`.`refueling_record` ;

CREATE TABLE IF NOT EXISTS `car-project`.`refueling_record` (
  `id` MEDIUMINT UNSIGNED NOT NULL COMMENT '加油紀錄編號',
  `vehicle_id` SMALLINT NOT NULL COMMENT '車輛ID',
  `date` VARCHAR(10) NOT NULL COMMENT '加油日期',
  `station` VARCHAR(45) NOT NULL COMMENT '加油站點',
  `octane_number` INT(4) UNSIGNED NOT NULL COMMENT '辛烷值',
  `unit_price` DECIMAL(3,1) UNSIGNED NULL DEFAULT 0 COMMENT '單價',
  `count` DECIMAL(4,2) UNSIGNED NULL DEFAULT 0 COMMENT '添加油量',
  `value` INT(6) UNSIGNED NOT NULL COMMENT '金額',
  `mileage` DECIMAL(7,2) UNSIGNED NOT NULL COMMENT '里程',
  `monitor_fuel_record` DECIMAL(3,1) UNSIGNED NULL DEFAULT 0 COMMENT '儀表油耗',
  PRIMARY KEY (`id`, `vehicle_id`),
  INDEX `INDEX_Vehicle_ID` USING BTREE (`vehicle_id`))
ENGINE = InnoDB
COMMENT = '加油紀錄';


-- -----------------------------------------------------
-- Table `car-project`.`maintenance_record`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `car-project`.`maintenance_record` ;

CREATE TABLE IF NOT EXISTS `car-project`.`maintenance_record` (
  `id` MEDIUMINT UNSIGNED NOT NULL COMMENT '維修編號',
  `vehicle_id` SMALLINT NOT NULL COMMENT '車輛ID',
  `date` VARCHAR(10) NOT NULL COMMENT '維修日期',
  `mileage` DECIMAL(7,2) UNSIGNED NULL DEFAULT 0 COMMENT '里程',
  `shop` VARCHAR(45) NOT NULL COMMENT '維修店家',
  `amount` INT(7) UNSIGNED NOT NULL COMMENT '金額',
  PRIMARY KEY (`id`, `vehicle_id`))
ENGINE = InnoDB
COMMENT = '加油紀錄';


-- -----------------------------------------------------
-- Table `car-project`.`maintenance_detail`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `car-project`.`maintenance_detail` ;

CREATE TABLE IF NOT EXISTS `car-project`.`maintenance_detail` (
  `id` TINYINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '維修項目ID',
  `maintenance_record_id` MEDIUMINT UNSIGNED NOT NULL COMMENT '維修紀錄ID',
  `name` VARCHAR(45) NULL DEFAULT '' COMMENT '維修品項',
  `value` INT(10) UNSIGNED NULL DEFAULT 0 COMMENT '單價',
  `content` VARCHAR(150) NULL DEFAULT '' COMMENT '內容, 料號, 工資, 備註',
  PRIMARY KEY (`id`, `maintenance_record_id`))
ENGINE = InnoDB
COMMENT = '維修項目';


-- -----------------------------------------------------
-- Table `car-project`.`user_gas_station`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `car-project`.`user_gas_station` ;

CREATE TABLE IF NOT EXISTS `car-project`.`user_gas_station` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` SMALLINT UNSIGNED NOT NULL,
  `name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`, `user_id`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
