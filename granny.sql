CREATE TABLE IF NOT EXISTS `tblContact` (
  `pkId` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `fldName` text COLLATE utf8_unicode_ci NOT NULL,
  `fldAddress` text COLLATE utf8_unicode_ci NOT NULL,
  `fldPhone` text COLLATE utf8_unicode_ci NOT NULL,
  `fldEmail` text COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`pkId`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='Table containing a list of Granny''s Contacts' AUTO_INCREMENT=1 ;
