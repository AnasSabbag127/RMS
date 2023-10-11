package api
//subadmin Routes
/*
	createUserHandler(w,r) {
		sql = `INSERT INTO user(name,email,password,role_id,is_admin,created_by) VALUES($1,$2,$3,$4,$5) `
		--- role_id == 1,2,3 represent respectfully admin,subadmin,user
	}	
	createDishesHandler(w,r){
		sql = `INSERT INTO dishes(restraunt_id,created_by,dish_name,price) VALUES($1,$2,$3.$4)`
	}
	createRestrauntHandler(w,r){
		sql = `INSERT INTO restraunt(name,address,created_by,stars) VALUES($1,$2,$3.$4)`
	}
	// subadmin can access only those user,restraunt,dish which is created by him
	getAllUserHandler(w,r){
		sql = `SELECT id,name,address,role_id FROM users WHERE created_by = $1`
	}
	getAllRestraunts(w,r){
		sql = `SELECT id,name,address,created_by,stars restraunt WHERE created_by = $1`
	}
	getAllDishes(w,r){
		sql = `SELECT id,restraunt_id,created_by,dish_name,price FROM dishes WHERE created_by = $1`
	}
	
*/