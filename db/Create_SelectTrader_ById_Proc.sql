USE [Markets]
GO
/****** Object:  StoredProcedure [dbo].[Select_Traders]    Script Date: 7/16/2018 2:25:34 PM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
-- =============================================
-- Author:		<Author,,Name>
-- Create date: <Create Date,,>
-- Description:	<Description,,>
-- =============================================
CREATE PROCEDURE [dbo].[Select_Trader_By_Id]
	@Id int
AS
BEGIN
	-- SET NOCOUNT ON added to prevent extra result sets from
	-- interfering with SELECT statements.
	SET NOCOUNT ON;

    -- Insert statements for procedure here
	SELECT trader.TraderID, trader.UserName, trader.Password, trader.Email, trader.IsVerified
	FROM Traders trader
	WHERE trader.TraderID = @Id
END
