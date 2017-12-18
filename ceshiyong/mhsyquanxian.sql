-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.5.35-MariaDB - mariadb.org binary distribution
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  8.0.0.4396
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
-- 正在导出表  mhsyquanxian.biaoji10s 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `biaoji10s` DISABLE KEYS */;
INSERT INTO `biaoji10s` (`id`, `biaobianma`, `bianma`, `fubianma`, `biaobianma7`, `zhi`, `miaoshu`) VALUES
	(1, 'bj1', 'bj1', 'zj', 'zd9', 'YOUXIAO:有效', '有效性的值范围'),
	(2, 'bj2', 'bj2', 'zj', 'zd9', 'WUXIAO:无效', '有效性的值范围');
/*!40000 ALTER TABLE `biaoji10s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.gl1he2s 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `gl1he2s` DISABLE KEYS */;
INSERT INTO `gl1he2s` (`id`, `biaobianma1`, `biaobianma2`) VALUES
	(1, 'zh3', 'zhl1');
/*!40000 ALTER TABLE `gl1he2s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.gl1he5s 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `gl1he5s` DISABLE KEYS */;
INSERT INTO `gl1he5s` (`id`, `biaobianma1`, `biaobianma5`) VALUES
	(1, 'zh3', 'js1');
/*!40000 ALTER TABLE `gl1he5s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.gl1he8s 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `gl1he8s` DISABLE KEYS */;
INSERT INTO `gl1he8s` (`id`, `biaobianma1`, `biaobianma8`) VALUES
	(1, 'zh3', 'lp1');
/*!40000 ALTER TABLE `gl1he8s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.gl2he4s 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `gl2he4s` DISABLE KEYS */;
INSERT INTO `gl2he4s` (`id`, `biaobianma2`, `biaobianma4`) VALUES
	(1, 'zhl1', 'yz3');
/*!40000 ALTER TABLE `gl2he4s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.gl5he6s 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `gl5he6s` DISABLE KEYS */;
INSERT INTO `gl5he6s` (`id`, `biaobianma5`, `biaobianma6`) VALUES
	(1, 'js1', 'zy1');
/*!40000 ALTER TABLE `gl5he6s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.gl8he9s 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `gl8he9s` DISABLE KEYS */;
INSERT INTO `gl8he9s` (`id`, `biaobianma8`, `biaobianma9`) VALUES
	(1, 'lp1', 'lpl1');
/*!40000 ALTER TABLE `gl8he9s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.juese5s 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `juese5s` DISABLE KEYS */;
INSERT INTO `juese5s` (`id`, `biaobianma`, `bianma`, `mingcheng`, `miaoshu`) VALUES
	(1, 'js1', 'GUANLIYUAN_XT', '系统管理员', '系统管理员');
/*!40000 ALTER TABLE `juese5s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.lingpai8s 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `lingpai8s` DISABLE KEYS */;
INSERT INTO `lingpai8s` (`id`, `biaobianma`, `bianma`, `zhi`, `biaoji`, `shijian`) VALUES
	(1, 'lp1', 'laizizh3', '123456', 'YOUXIAO,', '1970-01-01 10:00:00');
/*!40000 ALTER TABLE `lingpai8s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.lingpailei9s 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `lingpailei9s` DISABLE KEYS */;
INSERT INTO `lingpailei9s` (`id`, `biaobianma`, `bianma`, `mingcheng`, `guize`, `miaoshu`) VALUES
	(1, 'lpl1', 'JIAMI_YAN', '加密盐', 'JIAMI_YAN', '用盐的方式加密');
/*!40000 ALTER TABLE `lingpailei9s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.yanzheng4s 的数据：~3 rows (大约)
/*!40000 ALTER TABLE `yanzheng4s` DISABLE KEYS */;
INSERT INTO `yanzheng4s` (`id`, `biaobianma`, `bianma`, `mingcheng`, `jiami`, `fangfa`) VALUES
	(1, 'yz1', 'Youjian', '邮件', 'md5', 'Youjianjiaoyan'),
	(2, 'yz2', 'Shouji', '手机', 'md5', 'Duanxinjiaoyan'),
	(3, 'yz3', 'Yonghu', '密码', 'md5', 'Mimajiaoyan');
/*!40000 ALTER TABLE `yanzheng4s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.zhanghao1s 的数据：~3 rows (大约)
/*!40000 ALTER TABLE `zhanghao1s` DISABLE KEYS */;
INSERT INTO `zhanghao1s` (`id`, `biaobianma`, `bianma`, `shouquan`, `zhi`) VALUES
	(1, 'zh1', 'clz', 1, 'changlizhi@aliyun.com'),
	(2, 'zh2', 'clz', 1, '182184578455'),
	(3, 'zh3', 'clz', 1, '鱼饵');
/*!40000 ALTER TABLE `zhanghao1s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.zhanghaolei2s 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `zhanghaolei2s` DISABLE KEYS */;
INSERT INTO `zhanghaolei2s` (`id`, `biaobianma`, `bianma`, `mingcheng`) VALUES
	(1, 'zhl1', 'yonghuming', '用户名'),
	(2, 'zhl2', 'youxiang', '邮箱'),
	(3, 'zhl3', 'duanxin', '短信'),
	(4, 'zhl4', 'openid', 'openid');
/*!40000 ALTER TABLE `zhanghaolei2s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.zhanghaoxinxi3s 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `zhanghaoxinxi3s` DISABLE KEYS */;
INSERT INTO `zhanghaoxinxi3s` (`id`, `biaobianma1`, `biaobianma7`, `zhi`) VALUES
	(1, 'zh1', 'zd1', '60'),
	(2, 'zh1', 'zd1', '165');
/*!40000 ALTER TABLE `zhanghaoxinxi3s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.ziduan7s 的数据：~9 rows (大约)
/*!40000 ALTER TABLE `ziduan7s` DISABLE KEYS */;
INSERT INTO `ziduan7s` (`id`, `biaobianma`, `bianma`, `mingcheng`, `leixing`, `changdu`) VALUES
	(1, 'zd1', 'Nianling', '年龄', 'int', 3),
	(2, 'zd2', 'Shengao', '身高', 'int', 3),
	(3, 'zd3', 'Xuetang', '空腹血糖', 'float', 5),
	(4, 'zd4', 'Xueya', '血压', 'float', 5),
	(5, 'zd5', 'Zaoxuetang', '早餐后二小时血糖', 'float', 5),
	(6, 'zd6', 'Wuxuetang', '午餐后二小时血糖', 'float', 5),
	(7, 'zd7', 'Wanxuetang', '晚餐后二小时血糖', 'float', 5),
	(8, 'zd8', 'Gaomiduzhifangsuan', '高密度脂肪酸', 'float', 5),
	(9, 'zd9', 'Youxiaoxing', '有效性', 'string', 20);
/*!40000 ALTER TABLE `ziduan7s` ENABLE KEYS */;

-- 正在导出表  mhsyquanxian.ziyuan6s 的数据：~1 rows (大约)
/*!40000 ALTER TABLE `ziyuan6s` DISABLE KEYS */;
INSERT INTO `ziyuan6s` (`id`, `biaobianma`, `bianma`, `biaoshuzi`, `mingcheng`, `lujing`, `fangfa`, `miaoshu`) VALUES
	(1, 'zy1', 'ziduan_liebiao', 1, '字段列表', '/ziduan/liebiao', 'ziduanliebiao()', '获取所有字段的列表');
/*!40000 ALTER TABLE `ziyuan6s` ENABLE KEYS */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
