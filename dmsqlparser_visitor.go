// Code generated from DmSqlParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // DmSqlParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by DmSqlParser.
type DmSqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DmSqlParser#dmprogram.
	VisitDmprogram(ctx *DmprogramContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sql_clauses.
	VisitSql_clauses(ctx *Sql_clausesContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ddlsql.
	VisitDdlsql(ctx *DdlsqlContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dmlsql.
	VisitDmlsql(ctx *DmlsqlContext) interface{}

	// Visit a parse tree produced by DmSqlParser#privsql.
	VisitPrivsql(ctx *PrivsqlContext) interface{}

	// Visit a parse tree produced by DmSqlParser#othersql.
	VisitOthersql(ctx *OthersqlContext) interface{}

	// Visit a parse tree produced by DmSqlParser#utilsql.
	VisitUtilsql(ctx *UtilsqlContext) interface{}

	// Visit a parse tree produced by DmSqlParser#explainsql.
	VisitExplainsql(ctx *ExplainsqlContext) interface{}

	// Visit a parse tree produced by DmSqlParser#shutdown_stmt.
	VisitShutdown_stmt(ctx *Shutdown_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_diskgroup_stmt.
	VisitAlter_diskgroup_stmt(ctx *Alter_diskgroup_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#local.
	VisitLocal(ctx *LocalContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dmsubprogram.
	VisitDmsubprogram(ctx *DmsubprogramContext) interface{}

	// Visit a parse tree produced by DmSqlParser#declare_block.
	VisitDeclare_block(ctx *Declare_blockContext) interface{}

	// Visit a parse tree produced by DmSqlParser#decl_var_cur_list_options.
	VisitDecl_var_cur_list_options(ctx *Decl_var_cur_list_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#decl_var_cur_list_1.
	VisitDecl_var_cur_list_1(ctx *Decl_var_cur_list_1Context) interface{}

	// Visit a parse tree produced by DmSqlParser#decl_var_cur_list.
	VisitDecl_var_cur_list(ctx *Decl_var_cur_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#decl_plsql_type.
	VisitDecl_plsql_type(ctx *Decl_plsql_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#plsql_type_def.
	VisitPlsql_type_def(ctx *Plsql_type_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#lt_int_lst.
	VisitLt_int_lst(ctx *Lt_int_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#rec_item_def_list.
	VisitRec_item_def_list(ctx *Rec_item_def_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#rec_item_def.
	VisitRec_item_def(ctx *Rec_item_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#decl_variable.
	VisitDecl_variable(ctx *Decl_variableContext) interface{}

	// Visit a parse tree produced by DmSqlParser#not_null.
	VisitNot_null(ctx *Not_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#plsql_datatype.
	VisitPlsql_datatype(ctx *Plsql_datatypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#default_clause_option.
	VisitDefault_clause_option(ctx *Default_clause_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#variable_name_list.
	VisitVariable_name_list(ctx *Variable_name_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#decl_except.
	VisitDecl_except(ctx *Decl_exceptContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pragma_def.
	VisitPragma_def(ctx *Pragma_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pragma.
	VisitPragma(ctx *PragmaContext) interface{}

	// Visit a parse tree produced by DmSqlParser#plbody.
	VisitPlbody(ctx *PlbodyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ss_plbody.
	VisitSs_plbody(ctx *Ss_plbodyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by DmSqlParser#label_list.
	VisitLabel_list(ctx *Label_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#label_list_options.
	VisitLabel_list_options(ctx *Label_list_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#label_demiliter_l.
	VisitLabel_demiliter_l(ctx *Label_demiliter_lContext) interface{}

	// Visit a parse tree produced by DmSqlParser#label_demiliter_r.
	VisitLabel_demiliter_r(ctx *Label_demiliter_rContext) interface{}

	// Visit a parse tree produced by DmSqlParser#plsql.
	VisitPlsql(ctx *PlsqlContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ur_option.
	VisitUr_option(ctx *Ur_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#flashback_trig_enable.
	VisitFlashback_trig_enable(ctx *Flashback_trig_enableContext) interface{}

	// Visit a parse tree produced by DmSqlParser#scn_or_lsn.
	VisitScn_or_lsn(ctx *Scn_or_lsnContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_table_name_list.
	VisitFull_table_name_list(ctx *Full_table_name_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#flashback_tab_stmt.
	VisitFlashback_tab_stmt(ctx *Flashback_tab_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#rename.
	VisitRename(ctx *RenameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_system_set_stmt.
	VisitAlter_system_set_stmt(ctx *Alter_system_set_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#defer.
	VisitDefer(ctx *DeferContext) interface{}

	// Visit a parse tree produced by DmSqlParser#scope.
	VisitScope(ctx *ScopeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_session_stmt.
	VisitAlter_session_stmt(ctx *Alter_session_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#parallel_mode.
	VisitParallel_mode(ctx *Parallel_modeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#parallel_degree.
	VisitParallel_degree(ctx *Parallel_degreeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#purge.
	VisitPurge(ctx *PurgeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sess_id.
	VisitSess_id(ctx *Sess_idContext) interface{}

	// Visit a parse tree produced by DmSqlParser#set_time_zone_string.
	VisitSet_time_zone_string(ctx *Set_time_zone_stringContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sess_attr.
	VisitSess_attr(ctx *Sess_attrContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sess_attr_val.
	VisitSess_attr_val(ctx *Sess_attr_valContext) interface{}

	// Visit a parse tree produced by DmSqlParser#set_schema_stmt.
	VisitSet_schema_stmt(ctx *Set_schema_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#plblock.
	VisitPlblock(ctx *PlblockContext) interface{}

	// Visit a parse tree produced by DmSqlParser#except_option.
	VisitExcept_option(ctx *Except_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#finally_option.
	VisitFinally_option(ctx *Finally_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#finally_tail.
	VisitFinally_tail(ctx *Finally_tailContext) interface{}

	// Visit a parse tree produced by DmSqlParser#except_handler_list.
	VisitExcept_handler_list(ctx *Except_handler_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#except_handler.
	VisitExcept_handler(ctx *Except_handlerContext) interface{}

	// Visit a parse tree produced by DmSqlParser#except_name.
	VisitExcept_name(ctx *Except_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#except_list.
	VisitExcept_list(ctx *Except_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#if_stmt.
	VisitIf_stmt(ctx *If_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#if_stmt_clause.
	VisitIf_stmt_clause(ctx *If_stmt_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#if_condition_clause.
	VisitIf_condition_clause(ctx *If_condition_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#if_then_clause.
	VisitIf_then_clause(ctx *If_then_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#elseif_lst_option.
	VisitElseif_lst_option(ctx *Elseif_lst_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#elseif_clause.
	VisitElseif_clause(ctx *Elseif_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#else_option.
	VisitElse_option(ctx *Else_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ss_if_stmt_clause.
	VisitSs_if_stmt_clause(ctx *Ss_if_stmt_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ss_elseif_lst_option.
	VisitSs_elseif_lst_option(ctx *Ss_elseif_lst_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ss_elseif_clause.
	VisitSs_elseif_clause(ctx *Ss_elseif_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ss_else_option.
	VisitSs_else_option(ctx *Ss_else_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#case_stmt.
	VisitCase_stmt(ctx *Case_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#plsearched_when_clause.
	VisitPlsearched_when_clause(ctx *Plsearched_when_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#plsearched_when_list.
	VisitPlsearched_when_list(ctx *Plsearched_when_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#case_option.
	VisitCase_option(ctx *Case_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#assign_stmt.
	VisitAssign_stmt(ctx *Assign_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#assign_obj.
	VisitAssign_obj(ctx *Assign_objContext) interface{}

	// Visit a parse tree produced by DmSqlParser#assign_obj2.
	VisitAssign_obj2(ctx *Assign_obj2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#assign_op.
	VisitAssign_op(ctx *Assign_opContext) interface{}

	// Visit a parse tree produced by DmSqlParser#goto_stmt.
	VisitGoto_stmt(ctx *Goto_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#while_stmt.
	VisitWhile_stmt(ctx *While_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#loop_stmt.
	VisitLoop_stmt(ctx *Loop_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#repeat_stmt.
	VisitRepeat_stmt(ctx *Repeat_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#for_stmt.
	VisitFor_stmt(ctx *For_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#forall_stmt.
	VisitForall_stmt(ctx *Forall_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#forall_between_option.
	VisitForall_between_option(ctx *Forall_between_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#forall_save_exception_option.
	VisitForall_save_exception_option(ctx *Forall_save_exception_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#forall_index_values.
	VisitForall_index_values(ctx *Forall_index_valuesContext) interface{}

	// Visit a parse tree produced by DmSqlParser#forall_start.
	VisitForall_start(ctx *Forall_startContext) interface{}

	// Visit a parse tree produced by DmSqlParser#forall_dml_stmt.
	VisitForall_dml_stmt(ctx *Forall_dml_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#in_option.
	VisitIn_option(ctx *In_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#for_condition.
	VisitFor_condition(ctx *For_conditionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pipe_row_stmt.
	VisitPipe_row_stmt(ctx *Pipe_row_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#exit_stmt.
	VisitExit_stmt(ctx *Exit_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#continue_stmt.
	VisitContinue_stmt(ctx *Continue_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#null_stmt.
	VisitNull_stmt(ctx *Null_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#print_stmt.
	VisitPrint_stmt(ctx *Print_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#execute_stmt.
	VisitExecute_stmt(ctx *Execute_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dyn_return.
	VisitDyn_return(ctx *Dyn_returnContext) interface{}

	// Visit a parse tree produced by DmSqlParser#using_clause.
	VisitUsing_clause(ctx *Using_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#using_exp_list.
	VisitUsing_exp_list(ctx *Using_exp_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#using_exp.
	VisitUsing_exp(ctx *Using_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_proc_stmt.
	VisitAlter_proc_stmt(ctx *Alter_proc_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_func_stmt.
	VisitAlter_func_stmt(ctx *Alter_func_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_package_stmt.
	VisitAlter_package_stmt(ctx *Alter_package_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pkg_type.
	VisitPkg_type(ctx *Pkg_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#declare_opt.
	VisitDeclare_opt(ctx *Declare_optContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_table_stmt.
	VisitAlter_table_stmt(ctx *Alter_table_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_tag.
	VisitAlter_tag(ctx *Alter_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_index_stmt.
	VisitAlter_index_stmt(ctx *Alter_index_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_index_name.
	VisitFull_index_name(ctx *Full_index_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_index_action.
	VisitAlter_index_action(ctx *Alter_index_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#rebuild_clause.
	VisitRebuild_clause(ctx *Rebuild_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#exclusive_options.
	VisitExclusive_options(ctx *Exclusive_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#asynchronous_options.
	VisitAsynchronous_options(ctx *Asynchronous_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#visible_clause.
	VisitVisible_clause(ctx *Visible_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_def_list.
	VisitColumn_def_list(ctx *Column_def_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#lock.
	VisitLock(ctx *LockContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_table_partition_action.
	VisitAlter_table_partition_action(ctx *Alter_table_partition_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#template_info.
	VisitTemplate_info(ctx *Template_infoContext) interface{}

	// Visit a parse tree produced by DmSqlParser#template_item_2.
	VisitTemplate_item_2(ctx *Template_item_2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#template_item_1.
	VisitTemplate_item_1(ctx *Template_item_1Context) interface{}

	// Visit a parse tree produced by DmSqlParser#including_indexes.
	VisitIncluding_indexes(ctx *Including_indexesContext) interface{}

	// Visit a parse tree produced by DmSqlParser#truncate_partition_name.
	VisitTruncate_partition_name(ctx *Truncate_partition_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cons_enable.
	VisitCons_enable(ctx *Cons_enableContext) interface{}

	// Visit a parse tree produced by DmSqlParser#reuse_storage_option.
	VisitReuse_storage_option(ctx *Reuse_storage_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_table_action.
	VisitAlter_table_action(ctx *Alter_table_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#fast_flag.
	VisitFast_flag(ctx *Fast_flagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#storage_stat_flag.
	VisitStorage_stat_flag(ctx *Storage_stat_flagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#storage_stat_cols.
	VisitStorage_stat_cols(ctx *Storage_stat_colsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hfs_rebuild_level.
	VisitHfs_rebuild_level(ctx *Hfs_rebuild_levelContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ata_lock_option.
	VisitAta_lock_option(ctx *Ata_lock_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mdf_column_def_list.
	VisitMdf_column_def_list(ctx *Mdf_column_def_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mdf_column_def.
	VisitMdf_column_def(ctx *Mdf_column_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_def.
	VisitColumn_def(ctx *Column_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_def_ex.
	VisitColumn_def_ex(ctx *Column_def_exContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_def_low.
	VisitColumn_def_low(ctx *Column_def_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#virtual_column_datatype.
	VisitVirtual_column_datatype(ctx *Virtual_column_datatypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#virtual_column_generated.
	VisitVirtual_column_generated(ctx *Virtual_column_generatedContext) interface{}

	// Visit a parse tree produced by DmSqlParser#virtual_column_virtual.
	VisitVirtual_column_virtual(ctx *Virtual_column_virtualContext) interface{}

	// Visit a parse tree produced by DmSqlParser#virtual_column_visible.
	VisitVirtual_column_visible(ctx *Virtual_column_visibleContext) interface{}

	// Visit a parse tree produced by DmSqlParser#virtual_column_def.
	VisitVirtual_column_def(ctx *Virtual_column_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#charset_option.
	VisitCharset_option(ctx *Charset_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_def_4_option.
	VisitColumn_def_4_option(ctx *Column_def_4_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#auto_update_clause.
	VisitAuto_update_clause(ctx *Auto_update_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#update_exp.
	VisitUpdate_exp(ctx *Update_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#identity_clause.
	VisitIdentity_clause(ctx *Identity_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#default_clause_with_on_null_opt.
	VisitDefault_clause_with_on_null_opt(ctx *Default_clause_with_on_null_optContext) interface{}

	// Visit a parse tree produced by DmSqlParser#default_clause.
	VisitDefault_clause(ctx *Default_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#default_exp.
	VisitDefault_exp(ctx *Default_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_constraint_def.
	VisitColumn_constraint_def(ctx *Column_constraint_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#constraint_name_def_options.
	VisitConstraint_name_def_options(ctx *Constraint_name_def_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#constraint_name_def.
	VisitConstraint_name_def(ctx *Constraint_name_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_constraints.
	VisitColumn_constraints(ctx *Column_constraintsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_constraint.
	VisitColumn_constraint(ctx *Column_constraintContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_constraint_action.
	VisitColumn_constraint_action(ctx *Column_constraint_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#not_null_spec.
	VisitNot_null_spec(ctx *Not_null_specContext) interface{}

	// Visit a parse tree produced by DmSqlParser#unique_spec.
	VisitUnique_spec(ctx *Unique_specContext) interface{}

	// Visit a parse tree produced by DmSqlParser#refs_spec.
	VisitRefs_spec(ctx *Refs_specContext) interface{}

	// Visit a parse tree produced by DmSqlParser#refs_spec_action.
	VisitRefs_spec_action(ctx *Refs_spec_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#foreign_key.
	VisitForeign_key(ctx *Foreign_keyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#refd_table_and_columns.
	VisitRefd_table_and_columns(ctx *Refd_table_and_columnsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ref_column_list.
	VisitRef_column_list(ctx *Ref_column_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_list.
	VisitColumn_list(ctx *Column_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_list2.
	VisitColumn_list2(ctx *Column_list2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#full_column_list.
	VisitFull_column_list(ctx *Full_column_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_list_list.
	VisitColumn_list_list(ctx *Column_list_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#drop_column_list.
	VisitDrop_column_list(ctx *Drop_column_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#match_option.
	VisitMatch_option(ctx *Match_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#match_type.
	VisitMatch_type(ctx *Match_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ref_triggered_action.
	VisitRef_triggered_action(ctx *Ref_triggered_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#update_rule.
	VisitUpdate_rule(ctx *Update_ruleContext) interface{}

	// Visit a parse tree produced by DmSqlParser#delete_rule.
	VisitDelete_rule(ctx *Delete_ruleContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ref_action.
	VisitRef_action(ctx *Ref_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#check_constraint_def.
	VisitCheck_constraint_def(ctx *Check_constraint_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#check_condition.
	VisitCheck_condition(ctx *Check_conditionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restrict_cascade.
	VisitRestrict_cascade(ctx *Restrict_cascadeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cascade_opt.
	VisitCascade_opt(ctx *Cascade_optContext) interface{}

	// Visit a parse tree produced by DmSqlParser#constraint_name_options.
	VisitConstraint_name_options(ctx *Constraint_name_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#check_option_def_true.
	VisitCheck_option_def_true(ctx *Check_option_def_trueContext) interface{}

	// Visit a parse tree produced by DmSqlParser#constraint_attributes_options.
	VisitConstraint_attributes_options(ctx *Constraint_attributes_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#constraint_attributes.
	VisitConstraint_attributes(ctx *Constraint_attributesContext) interface{}

	// Visit a parse tree produced by DmSqlParser#deferrable_option.
	VisitDeferrable_option(ctx *Deferrable_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#constraint_check_time.
	VisitConstraint_check_time(ctx *Constraint_check_timeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#table_constraint_clause.
	VisitTable_constraint_clause(ctx *Table_constraint_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#table_constraint.
	VisitTable_constraint(ctx *Table_constraintContext) interface{}

	// Visit a parse tree produced by DmSqlParser#using_index_clause.
	VisitUsing_index_clause(ctx *Using_index_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#foreign_key_clause.
	VisitForeign_key_clause(ctx *Foreign_key_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_trigger_stmt.
	VisitAlter_trigger_stmt(ctx *Alter_trigger_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_trigger_option.
	VisitAlter_trigger_option(ctx *Alter_trigger_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_table_partition_action_options.
	VisitAlter_table_partition_action_options(ctx *Alter_table_partition_action_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#refresh_materialized_view_stmt.
	VisitRefresh_materialized_view_stmt(ctx *Refresh_materialized_view_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#complete_del_null.
	VisitComplete_del_null(ctx *Complete_del_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#refresh_complete_del.
	VisitRefresh_complete_del(ctx *Refresh_complete_delContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_materialized_view_stmt.
	VisitAlter_materialized_view_stmt(ctx *Alter_materialized_view_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_view_stmt.
	VisitAlter_view_stmt(ctx *Alter_view_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_view_action.
	VisitAlter_view_action(ctx *Alter_view_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cons_novalidate.
	VisitCons_novalidate(ctx *Cons_novalidateContext) interface{}

	// Visit a parse tree produced by DmSqlParser#view_constraint_clause.
	VisitView_constraint_clause(ctx *View_constraint_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#view_constraint.
	VisitView_constraint(ctx *View_constraintContext) interface{}

	// Visit a parse tree produced by DmSqlParser#view_unique_spec.
	VisitView_unique_spec(ctx *View_unique_specContext) interface{}

	// Visit a parse tree produced by DmSqlParser#view_refs_spec.
	VisitView_refs_spec(ctx *View_refs_specContext) interface{}

	// Visit a parse tree produced by DmSqlParser#view_refs_spec_action.
	VisitView_refs_spec_action(ctx *View_refs_spec_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#call_proc_stmt.
	VisitCall_proc_stmt(ctx *Call_proc_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#raw_call_proc_stmt.
	VisitRaw_call_proc_stmt(ctx *Raw_call_proc_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#call_proc_stmt_2.
	VisitCall_proc_stmt_2(ctx *Call_proc_stmt_2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#exec_proc_stmt.
	VisitExec_proc_stmt(ctx *Exec_proc_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dblink_clause.
	VisitDblink_clause(ctx *Dblink_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dblink_clause2.
	VisitDblink_clause2(ctx *Dblink_clause2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#param_list.
	VisitParam_list(ctx *Param_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by DmSqlParser#raw_exp_list.
	VisitRaw_exp_list(ctx *Raw_exp_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#exp_list_2.
	VisitExp_list_2(ctx *Exp_list_2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#exp_list.
	VisitExp_list(ctx *Exp_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ins_exp_list.
	VisitIns_exp_list(ctx *Ins_exp_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#lt_exp.
	VisitLt_exp(ctx *Lt_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#range_partition_exp.
	VisitRange_partition_exp(ctx *Range_partition_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#range_partition_exp_list.
	VisitRange_partition_exp_list(ctx *Range_partition_exp_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#list_partition_exp.
	VisitList_partition_exp(ctx *List_partition_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#list_partition_exp_list.
	VisitList_partition_exp_list(ctx *List_partition_exp_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#list_partition_value_list.
	VisitList_partition_value_list(ctx *List_partition_value_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#close_cursor_stmt.
	VisitClose_cursor_stmt(ctx *Close_cursor_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#close_cursor_statement.
	VisitClose_cursor_statement(ctx *Close_cursor_statementContext) interface{}

	// Visit a parse tree produced by DmSqlParser#begin_trans_stmt.
	VisitBegin_trans_stmt(ctx *Begin_trans_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#commit_trans_stmt.
	VisitCommit_trans_stmt(ctx *Commit_trans_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#commit_head.
	VisitCommit_head(ctx *Commit_headContext) interface{}

	// Visit a parse tree produced by DmSqlParser#commit_tail.
	VisitCommit_tail(ctx *Commit_tailContext) interface{}

	// Visit a parse tree produced by DmSqlParser#commit_wait_immed_option.
	VisitCommit_wait_immed_option(ctx *Commit_wait_immed_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#connect_stmt.
	VisitConnect_stmt(ctx *Connect_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#password.
	VisitPassword(ctx *PasswordContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ts_storage.
	VisitTs_storage(ctx *Ts_storageContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ts_storage_clause.
	VisitTs_storage_clause(ctx *Ts_storage_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_tablespace_stmt.
	VisitCreate_tablespace_stmt(ctx *Create_tablespace_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ctss_with_clause.
	VisitCtss_with_clause(ctx *Ctss_with_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_tablespace_set_stmt.
	VisitCreate_tablespace_set_stmt(ctx *Create_tablespace_set_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_tablespace_set_stmt.
	VisitAlter_tablespace_set_stmt(ctx *Alter_tablespace_set_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cache.
	VisitCache(ctx *CacheContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_tablespace_stmt.
	VisitAlter_tablespace_stmt(ctx *Alter_tablespace_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#keep.
	VisitKeep(ctx *KeepContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_tablespace_action.
	VisitAlter_tablespace_action(ctx *Alter_tablespace_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#file_list.
	VisitFile_list(ctx *File_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pathname_list.
	VisitPathname_list(ctx *Pathname_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#integer_list.
	VisitInteger_list(ctx *Integer_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mirror.
	VisitMirror(ctx *MirrorContext) interface{}

	// Visit a parse tree produced by DmSqlParser#autoextend_nextsize.
	VisitAutoextend_nextsize(ctx *Autoextend_nextsizeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#autoextend_maxsize.
	VisitAutoextend_maxsize(ctx *Autoextend_maxsizeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#autoextend.
	VisitAutoextend(ctx *AutoextendContext) interface{}

	// Visit a parse tree produced by DmSqlParser#on_raft.
	VisitOn_raft(ctx *On_raftContext) interface{}

	// Visit a parse tree produced by DmSqlParser#archfile.
	VisitArchfile(ctx *ArchfileContext) interface{}

	// Visit a parse tree produced by DmSqlParser#archflag.
	VisitArchflag(ctx *ArchflagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#archstyle_options.
	VisitArchstyle_options(ctx *Archstyle_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#archstyle.
	VisitArchstyle(ctx *ArchstyleContext) interface{}

	// Visit a parse tree produced by DmSqlParser#archdir.
	VisitArchdir(ctx *ArchdirContext) interface{}

	// Visit a parse tree produced by DmSqlParser#bakfile.
	VisitBakfile(ctx *BakfileContext) interface{}

	// Visit a parse tree produced by DmSqlParser#parameters_option_list.
	VisitParameters_option_list(ctx *Parameters_option_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#parameter_option_list.
	VisitParameter_option_list(ctx *Parameter_option_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#parameter_option.
	VisitParameter_option(ctx *Parameter_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pathname.
	VisitPathname(ctx *PathnameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pathname_options.
	VisitPathname_options(ctx *Pathname_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_stmt.
	VisitBackup_stmt(ctx *Backup_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#back_range_option.
	VisitBack_range_option(ctx *Back_range_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#back_archive_spec_null.
	VisitBack_archive_spec_null(ctx *Back_archive_spec_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#not_backed_up.
	VisitNot_backed_up(ctx *Not_backed_upContext) interface{}

	// Visit a parse tree produced by DmSqlParser#archive_spec.
	VisitArchive_spec(ctx *Archive_specContext) interface{}

	// Visit a parse tree produced by DmSqlParser#spec_lsn.
	VisitSpec_lsn(ctx *Spec_lsnContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_delete_archive.
	VisitBackup_delete_archive(ctx *Backup_delete_archiveContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_options.
	VisitBackup_options(ctx *Backup_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cumulative.
	VisitCumulative(ctx *CumulativeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_bak_dir_list.
	VisitWith_bak_dir_list(ctx *With_bak_dir_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#base_on_backup.
	VisitBase_on_backup(ctx *Base_on_backupContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_to_options.
	VisitBackup_to_options(ctx *Backup_to_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_path_null.
	VisitBackup_path_null(ctx *Backup_path_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#device_type.
	VisitDevice_type(ctx *Device_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#parms_command.
	VisitParms_command(ctx *Parms_commandContext) interface{}

	// Visit a parse tree produced by DmSqlParser#media_name.
	VisitMedia_name(ctx *Media_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_desc_options.
	VisitBackup_desc_options(ctx *Backup_desc_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_desc.
	VisitBackup_desc(ctx *Backup_descContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_maxsize.
	VisitBackup_maxsize(ctx *Backup_maxsizeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_limit.
	VisitBackup_limit(ctx *Backup_limitContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_identified.
	VisitBackup_identified(ctx *Backup_identifiedContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_compressed.
	VisitBackup_compressed(ctx *Backup_compressedContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_without.
	VisitBackup_without(ctx *Backup_withoutContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_tsk_thread_num_null.
	VisitBackup_tsk_thread_num_null(ctx *Backup_tsk_thread_num_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_parallel_dir.
	VisitBackup_parallel_dir(ctx *Backup_parallel_dirContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_trace_file_level.
	VisitBackup_trace_file_level(ctx *Backup_trace_file_levelContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restore_stmt.
	VisitRestore_stmt(ctx *Restore_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restore_datafile_lst.
	VisitRestore_datafile_lst(ctx *Restore_datafile_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restore_mapped_file.
	VisitRestore_mapped_file(ctx *Restore_mapped_fileContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mapped_file.
	VisitMapped_file(ctx *Mapped_fileContext) interface{}

	// Visit a parse tree produced by DmSqlParser#res_struct.
	VisitRes_struct(ctx *Res_structContext) interface{}

	// Visit a parse tree produced by DmSqlParser#tsk_thread_num.
	VisitTsk_thread_num(ctx *Tsk_thread_numContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restore_tsk_thread_num_null.
	VisitRestore_tsk_thread_num_null(ctx *Restore_tsk_thread_num_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restore_parallel.
	VisitRestore_parallel(ctx *Restore_parallelContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_table_name_options.
	VisitFull_table_name_options(ctx *Full_table_name_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#res_without_index_constraint.
	VisitRes_without_index_constraint(ctx *Res_without_index_constraintContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restore_from.
	VisitRestore_from(ctx *Restore_fromContext) interface{}

	// Visit a parse tree produced by DmSqlParser#res_until.
	VisitRes_until(ctx *Res_untilContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restore_file_list_options.
	VisitRestore_file_list_options(ctx *Restore_file_list_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mirror_file_list_options.
	VisitMirror_file_list_options(ctx *Mirror_file_list_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restore_trace_file_level.
	VisitRestore_trace_file_level(ctx *Restore_trace_file_levelContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restore_file_list.
	VisitRestore_file_list(ctx *Restore_file_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restore_file.
	VisitRestore_file(ctx *Restore_fileContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mirror_file_list.
	VisitMirror_file_list(ctx *Mirror_file_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mirror_file.
	VisitMirror_file(ctx *Mirror_fileContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_bak_arch_dir_list.
	VisitWith_bak_arch_dir_list(ctx *With_bak_arch_dir_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restore_identified.
	VisitRestore_identified(ctx *Restore_identifiedContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_func_stmt.
	VisitCreate_func_stmt(ctx *Create_func_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#func_aggr_clause.
	VisitFunc_aggr_clause(ctx *Func_aggr_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pipelined_options.
	VisitPipelined_options(ctx *Pipelined_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#replace_option.
	VisitReplace_option(ctx *Replace_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#edit_options.
	VisitEdit_options(ctx *Edit_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#encryption_option.
	VisitEncryption_option(ctx *Encryption_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#calc_option.
	VisitCalc_option(ctx *Calc_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#func_action.
	VisitFunc_action(ctx *Func_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#func_call_options.
	VisitFunc_call_options(ctx *Func_call_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#func_call_option_list.
	VisitFunc_call_option_list(ctx *Func_call_option_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#func_call_option.
	VisitFunc_call_option(ctx *Func_call_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#invoker_rights_clause_options.
	VisitInvoker_rights_clause_options(ctx *Invoker_rights_clause_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#invoker_rights_clause.
	VisitInvoker_rights_clause(ctx *Invoker_rights_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#deterministic_clause_options.
	VisitDeterministic_clause_options(ctx *Deterministic_clause_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#deterministic_clause.
	VisitDeterministic_clause(ctx *Deterministic_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#func_call_option2_options.
	VisitFunc_call_option2_options(ctx *Func_call_option2_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#func_call_option_list2.
	VisitFunc_call_option_list2(ctx *Func_call_option_list2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#func_call_option2.
	VisitFunc_call_option2(ctx *Func_call_option2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#result_cache_clause.
	VisitResult_cache_clause(ctx *Result_cache_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#inner_fun_name.
	VisitInner_fun_name(ctx *Inner_fun_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#platform_type.
	VisitPlatform_type(ctx *Platform_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#param_def_list_option.
	VisitParam_def_list_option(ctx *Param_def_list_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#param_def_list.
	VisitParam_def_list(ctx *Param_def_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#param_def.
	VisitParam_def(ctx *Param_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#param_in_out_option.
	VisitParam_in_out_option(ctx *Param_in_out_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#is_as.
	VisitIs_as(ctx *Is_asContext) interface{}

	// Visit a parse tree produced by DmSqlParser#stat_on_org_stmt.
	VisitStat_on_org_stmt(ctx *Stat_on_org_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#stat_size.
	VisitStat_size(ctx *Stat_sizeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#stat_para.
	VisitStat_para(ctx *Stat_paraContext) interface{}

	// Visit a parse tree produced by DmSqlParser#stat_summarize.
	VisitStat_summarize(ctx *Stat_summarizeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mstat_ex.
	VisitMstat_ex(ctx *Mstat_exContext) interface{}

	// Visit a parse tree produced by DmSqlParser#indexid.
	VisitIndexid(ctx *IndexidContext) interface{}

	// Visit a parse tree produced by DmSqlParser#global_tag.
	VisitGlobal_tag(ctx *Global_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#bm_join_index_clause.
	VisitBm_join_index_clause(ctx *Bm_join_index_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#parallel_stmt.
	VisitParallel_stmt(ctx *Parallel_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_index_stmt.
	VisitCreate_index_stmt(ctx *Create_index_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_inner.
	VisitWith_inner(ctx *With_innerContext) interface{}

	// Visit a parse tree produced by DmSqlParser#index_no_sort.
	VisitIndex_no_sort(ctx *Index_no_sortContext) interface{}

	// Visit a parse tree produced by DmSqlParser#online_options.
	VisitOnline_options(ctx *Online_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#unusable_options.
	VisitUnusable_options(ctx *Unusable_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#reverse_options.
	VisitReverse_options(ctx *Reverse_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#index_column_list.
	VisitIndex_column_list(ctx *Index_column_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#index_column_name.
	VisitIndex_column_name(ctx *Index_column_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#storage_hash_tag.
	VisitStorage_hash_tag(ctx *Storage_hash_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#storage_hash.
	VisitStorage_hash(ctx *Storage_hashContext) interface{}

	// Visit a parse tree produced by DmSqlParser#storage_tag.
	VisitStorage_tag(ctx *Storage_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#storage_tag_nn.
	VisitStorage_tag_nn(ctx *Storage_tag_nnContext) interface{}

	// Visit a parse tree produced by DmSqlParser#tablespace_clause.
	VisitTablespace_clause(ctx *Tablespace_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#object_table_substitution_clause.
	VisitObject_table_substitution_clause(ctx *Object_table_substitution_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#object_table_substitution.
	VisitObject_table_substitution(ctx *Object_table_substitutionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#oid_clause.
	VisitOid_clause(ctx *Oid_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#oid_gen_type.
	VisitOid_gen_type(ctx *Oid_gen_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#oid_index_clause.
	VisitOid_index_clause(ctx *Oid_index_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#oid_tablespace_clause.
	VisitOid_tablespace_clause(ctx *Oid_tablespace_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#oid_tablespace_name.
	VisitOid_tablespace_name(ctx *Oid_tablespace_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#local_option.
	VisitLocal_option(ctx *Local_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#storage_list.
	VisitStorage_list(ctx *Storage_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#storage_hashpartmap.
	VisitStorage_hashpartmap(ctx *Storage_hashpartmapContext) interface{}

	// Visit a parse tree produced by DmSqlParser#storage.
	VisitStorage(ctx *StorageContext) interface{}

	// Visit a parse tree produced by DmSqlParser#id_list.
	VisitId_list(ctx *Id_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_proc_stmt.
	VisitCreate_proc_stmt(ctx *Create_proc_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_package_stmt.
	VisitCreate_package_stmt(ctx *Create_package_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pkg_cls_flag.
	VisitPkg_cls_flag(ctx *Pkg_cls_flagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#blk_end_option.
	VisitBlk_end_option(ctx *Blk_end_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#blk_end_option_low.
	VisitBlk_end_option_low(ctx *Blk_end_option_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#package_def_list_options.
	VisitPackage_def_list_options(ctx *Package_def_list_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#package_def_list.
	VisitPackage_def_list(ctx *Package_def_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#package_def.
	VisitPackage_def(ctx *Package_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#restrict_param_lst.
	VisitRestrict_param_lst(ctx *Restrict_param_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_package_body_stmt.
	VisitCreate_package_body_stmt(ctx *Create_package_body_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_pkg_body_header.
	VisitCreate_pkg_body_header(ctx *Create_pkg_body_headerContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pkg_cls_body_flag.
	VisitPkg_cls_body_flag(ctx *Pkg_cls_body_flagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#package_body_init_option.
	VisitPackage_body_init_option(ctx *Package_body_init_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#package_body_def_list.
	VisitPackage_body_def_list(ctx *Package_body_def_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#package_body_def.
	VisitPackage_body_def(ctx *Package_body_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#package_body_def2.
	VisitPackage_body_def2(ctx *Package_body_def2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#check_exec_options.
	VisitCheck_exec_options(ctx *Check_exec_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpg_decl_stmt.
	VisitSubpg_decl_stmt(ctx *Subpg_decl_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_type_stmt.
	VisitCreate_type_stmt(ctx *Create_type_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#force_option.
	VisitForce_option(ctx *Force_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#object_option.
	VisitObject_option(ctx *Object_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#under_option.
	VisitUnder_option(ctx *Under_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#object_def_list_options.
	VisitObject_def_list_options(ctx *Object_def_list_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#object_def_list.
	VisitObject_def_list(ctx *Object_def_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#object_def.
	VisitObject_def(ctx *Object_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#member_static.
	VisitMember_static(ctx *Member_staticContext) interface{}

	// Visit a parse tree produced by DmSqlParser#method_inherit_options.
	VisitMethod_inherit_options(ctx *Method_inherit_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#method_inherit_option.
	VisitMethod_inherit_option(ctx *Method_inherit_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#final_inst_list_options.
	VisitFinal_inst_list_options(ctx *Final_inst_list_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#final_inst_list.
	VisitFinal_inst_list(ctx *Final_inst_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#final_inst.
	VisitFinal_inst(ctx *Final_instContext) interface{}

	// Visit a parse tree produced by DmSqlParser#overriding_option.
	VisitOverriding_option(ctx *Overriding_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#method_attr_options.
	VisitMethod_attr_options(ctx *Method_attr_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#method_attr.
	VisitMethod_attr(ctx *Method_attrContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_type_body_stmt.
	VisitCreate_type_body_stmt(ctx *Create_type_body_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#object_body_def_list.
	VisitObject_body_def_list(ctx *Object_body_def_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#object_body_def.
	VisitObject_body_def(ctx *Object_body_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_context_stmt.
	VisitCreate_context_stmt(ctx *Create_context_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#namespace.
	VisitNamespace(ctx *NamespaceContext) interface{}

	// Visit a parse tree produced by DmSqlParser#initialized.
	VisitInitialized(ctx *InitializedContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_directory_stmt.
	VisitCreate_directory_stmt(ctx *Create_directory_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_crypto_stmt.
	VisitCreate_crypto_stmt(ctx *Create_crypto_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_crypto_stmt.
	VisitAlter_crypto_stmt(ctx *Alter_crypto_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_crypto_action.
	VisitAlter_crypto_action(ctx *Alter_crypto_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#comment_stmt.
	VisitComment_stmt(ctx *Comment_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#comment_tag.
	VisitComment_tag(ctx *Comment_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_partition_group_stmt.
	VisitCreate_partition_group_stmt(ctx *Create_partition_group_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#storage_act_datatype.
	VisitStorage_act_datatype(ctx *Storage_act_datatypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pg_storage_lst.
	VisitPg_storage_lst(ctx *Pg_storage_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pg_storage.
	VisitPg_storage(ctx *Pg_storageContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_table_stmt.
	VisitCreate_table_stmt(ctx *Create_table_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ctab_append_attr_clause.
	VisitCtab_append_attr_clause(ctx *Ctab_append_attr_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ctab_append_attr_list.
	VisitCtab_append_attr_list(ctx *Ctab_append_attr_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cobjtab_append_attr_clause.
	VisitCobjtab_append_attr_clause(ctx *Cobjtab_append_attr_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cobjtab_append_attr_list.
	VisitCobjtab_append_attr_list(ctx *Cobjtab_append_attr_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ctab_append_attr.
	VisitCtab_append_attr(ctx *Ctab_append_attrContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cobjtab_append_attr.
	VisitCobjtab_append_attr(ctx *Cobjtab_append_attrContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_table_action.
	VisitCreate_table_action(ctx *Create_table_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ctab_log_options.
	VisitCtab_log_options(ctx *Ctab_log_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ctab_log_option.
	VisitCtab_log_option(ctx *Ctab_log_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ctab_error_options.
	VisitCtab_error_options(ctx *Ctab_error_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#advance_log_clause.
	VisitAdvance_log_clause(ctx *Advance_log_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#add_log_clause.
	VisitAdd_log_clause(ctx *Add_log_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ctab_error_option.
	VisitCtab_error_option(ctx *Ctab_error_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ctab_row_movement_clause.
	VisitCtab_row_movement_clause(ctx *Ctab_row_movement_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#range_distribute_act.
	VisitRange_distribute_act(ctx *Range_distribute_actContext) interface{}

	// Visit a parse tree produced by DmSqlParser#range_distribute_act_lst.
	VisitRange_distribute_act_lst(ctx *Range_distribute_act_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#list_distribute_act.
	VisitList_distribute_act(ctx *List_distribute_actContext) interface{}

	// Visit a parse tree produced by DmSqlParser#list_distribute_act_list.
	VisitList_distribute_act_list(ctx *List_distribute_act_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#distribute_by_option.
	VisitDistribute_by_option(ctx *Distribute_by_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#distribute_by.
	VisitDistribute_by(ctx *Distribute_byContext) interface{}

	// Visit a parse tree produced by DmSqlParser#increment_set.
	VisitIncrement_set(ctx *Increment_setContext) interface{}

	// Visit a parse tree produced by DmSqlParser#increment.
	VisitIncrement(ctx *IncrementContext) interface{}

	// Visit a parse tree produced by DmSqlParser#rowdependencies_clause.
	VisitRowdependencies_clause(ctx *Rowdependencies_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pg_sub_partition.
	VisitPg_sub_partition(ctx *Pg_sub_partitionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#table_type_option.
	VisitTable_type_option(ctx *Table_type_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#table_temp_option.
	VisitTable_temp_option(ctx *Table_temp_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#objtab_elem_constraint.
	VisitObjtab_elem_constraint(ctx *Objtab_elem_constraintContext) interface{}

	// Visit a parse tree produced by DmSqlParser#objtab_element_cons_list.
	VisitObjtab_element_cons_list(ctx *Objtab_element_cons_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#objtab_element_cons.
	VisitObjtab_element_cons(ctx *Objtab_element_consContext) interface{}

	// Visit a parse tree produced by DmSqlParser#objcol_constraint.
	VisitObjcol_constraint(ctx *Objcol_constraintContext) interface{}

	// Visit a parse tree produced by DmSqlParser#table_element_list.
	VisitTable_element_list(ctx *Table_element_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#table_element.
	VisitTable_element(ctx *Table_elementContext) interface{}

	// Visit a parse tree produced by DmSqlParser#table_constraint_def.
	VisitTable_constraint_def(ctx *Table_constraint_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#on_commit_option.
	VisitOn_commit_option(ctx *On_commit_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#on_commit_option_nn.
	VisitOn_commit_option_nn(ctx *On_commit_option_nnContext) interface{}

	// Visit a parse tree produced by DmSqlParser#logging_option.
	VisitLogging_option(ctx *Logging_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#logging_option_nn.
	VisitLogging_option_nn(ctx *Logging_option_nnContext) interface{}

	// Visit a parse tree produced by DmSqlParser#partition_clause.
	VisitPartition_clause(ctx *Partition_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#partition_clause_nn.
	VisitPartition_clause_nn(ctx *Partition_clause_nnContext) interface{}

	// Visit a parse tree produced by DmSqlParser#horizon_partition_clause.
	VisitHorizon_partition_clause(ctx *Horizon_partition_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#compress_tag_hdr.
	VisitCompress_tag_hdr(ctx *Compress_tag_hdrContext) interface{}

	// Visit a parse tree produced by DmSqlParser#compress_clause_opt.
	VisitCompress_clause_opt(ctx *Compress_clause_optContext) interface{}

	// Visit a parse tree produced by DmSqlParser#compress_tag.
	VisitCompress_tag(ctx *Compress_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#compress_level.
	VisitCompress_level(ctx *Compress_levelContext) interface{}

	// Visit a parse tree produced by DmSqlParser#compress_type.
	VisitCompress_type(ctx *Compress_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#range_partition.
	VisitRange_partition(ctx *Range_partitionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#range_partition_list.
	VisitRange_partition_list(ctx *Range_partition_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hash_partition.
	VisitHash_partition(ctx *Hash_partitionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hash_partition_list.
	VisitHash_partition_list(ctx *Hash_partition_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#list_partition.
	VisitList_partition(ctx *List_partitionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#list_partition_list.
	VisitList_partition_list(ctx *List_partition_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#split_partition_list.
	VisitSplit_partition_list(ctx *Split_partition_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#partition_act.
	VisitPartition_act(ctx *Partition_actContext) interface{}

	// Visit a parse tree produced by DmSqlParser#vertical_partition_act.
	VisitVertical_partition_act(ctx *Vertical_partition_actContext) interface{}

	// Visit a parse tree produced by DmSqlParser#interval_item.
	VisitInterval_item(ctx *Interval_itemContext) interface{}

	// Visit a parse tree produced by DmSqlParser#horizon_partition_act_datatype.
	VisitHorizon_partition_act_datatype(ctx *Horizon_partition_act_datatypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#horizon_partition_act.
	VisitHorizon_partition_act(ctx *Horizon_partition_actContext) interface{}

	// Visit a parse tree produced by DmSqlParser#lock_partitions_clause.
	VisitLock_partitions_clause(ctx *Lock_partitions_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartion_template_list_datatype_options.
	VisitSubpartion_template_list_datatype_options(ctx *Subpartion_template_list_datatype_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartion_template_list_datatype.
	VisitSubpartion_template_list_datatype(ctx *Subpartion_template_list_datatypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartion_template_list_options.
	VisitSubpartion_template_list_options(ctx *Subpartion_template_list_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartion_template_list.
	VisitSubpartion_template_list(ctx *Subpartion_template_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartion_template_datatype.
	VisitSubpartion_template_datatype(ctx *Subpartion_template_datatypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#range_subpartion_template_datatype.
	VisitRange_subpartion_template_datatype(ctx *Range_subpartion_template_datatypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hash_subpartion_template_datatype.
	VisitHash_subpartion_template_datatype(ctx *Hash_subpartion_template_datatypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hash_subpartions_template_datatype_option.
	VisitHash_subpartions_template_datatype_option(ctx *Hash_subpartions_template_datatype_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#list_subpartion_template_datatype.
	VisitList_subpartion_template_datatype(ctx *List_subpartion_template_datatypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartion_template.
	VisitSubpartion_template(ctx *Subpartion_templateContext) interface{}

	// Visit a parse tree produced by DmSqlParser#range_subpartion_template.
	VisitRange_subpartion_template(ctx *Range_subpartion_templateContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hash_subpartion_template.
	VisitHash_subpartion_template(ctx *Hash_subpartion_templateContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hash_subpartions_template_option.
	VisitHash_subpartions_template_option(ctx *Hash_subpartions_template_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#list_subpartion_template.
	VisitList_subpartion_template(ctx *List_subpartion_templateContext) interface{}

	// Visit a parse tree produced by DmSqlParser#range_subpartition.
	VisitRange_subpartition(ctx *Range_subpartitionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hash_subpartition.
	VisitHash_subpartition(ctx *Hash_subpartitionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#list_subpartition.
	VisitList_subpartition(ctx *List_subpartitionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#range_subpartition_list.
	VisitRange_subpartition_list(ctx *Range_subpartition_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hash_subpartition_list.
	VisitHash_subpartition_list(ctx *Hash_subpartition_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#list_subpartition_list.
	VisitList_subpartition_list(ctx *List_subpartition_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartition_hash_desc.
	VisitSubpartition_hash_desc(ctx *Subpartition_hash_descContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartition_range_desc.
	VisitSubpartition_range_desc(ctx *Subpartition_range_descContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartition_list_desc.
	VisitSubpartition_list_desc(ctx *Subpartition_list_descContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartition_hash_desc_list.
	VisitSubpartition_hash_desc_list(ctx *Subpartition_hash_desc_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartition_range_desc_list.
	VisitSubpartition_range_desc_list(ctx *Subpartition_range_desc_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartition_list_desc_list.
	VisitSubpartition_list_desc_list(ctx *Subpartition_list_desc_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartition_desc.
	VisitSubpartition_desc(ctx *Subpartition_descContext) interface{}

	// Visit a parse tree produced by DmSqlParser#subpartition_desc_option.
	VisitSubpartition_desc_option(ctx *Subpartition_desc_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#add_subpartition_desc.
	VisitAdd_subpartition_desc(ctx *Add_subpartition_descContext) interface{}

	// Visit a parse tree produced by DmSqlParser#partition_no.
	VisitPartition_no(ctx *Partition_noContext) interface{}

	// Visit a parse tree produced by DmSqlParser#comment_clause.
	VisitComment_clause(ctx *Comment_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#encrypt_clause_options.
	VisitEncrypt_clause_options(ctx *Encrypt_clause_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#encrypt_clause.
	VisitEncrypt_clause(ctx *Encrypt_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#encrypt_cipher.
	VisitEncrypt_cipher(ctx *Encrypt_cipherContext) interface{}

	// Visit a parse tree produced by DmSqlParser#crypto_name.
	VisitCrypto_name(ctx *Crypto_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cipher_name.
	VisitCipher_name(ctx *Cipher_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_cipher_name.
	VisitFull_cipher_name(ctx *Full_cipher_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#encrypt_type.
	VisitEncrypt_type(ctx *Encrypt_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#manual_clause.
	VisitManual_clause(ctx *Manual_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#user_clause_options.
	VisitUser_clause_options(ctx *User_clause_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#user_clause.
	VisitUser_clause(ctx *User_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#user_list_option.
	VisitUser_list_option(ctx *User_list_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#user_list.
	VisitUser_list(ctx *User_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hash_cipher.
	VisitHash_cipher(ctx *Hash_cipherContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hash_type.
	VisitHash_type(ctx *Hash_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#space_limit.
	VisitSpace_limit(ctx *Space_limitContext) interface{}

	// Visit a parse tree produced by DmSqlParser#space_limit_nn.
	VisitSpace_limit_nn(ctx *Space_limit_nnContext) interface{}

	// Visit a parse tree produced by DmSqlParser#space_limit_1.
	VisitSpace_limit_1(ctx *Space_limit_1Context) interface{}

	// Visit a parse tree produced by DmSqlParser#space_limit2.
	VisitSpace_limit2(ctx *Space_limit2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#del_res.
	VisitDel_res(ctx *Del_resContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_enable.
	VisitTrig_enable(ctx *Trig_enableContext) interface{}

	// Visit a parse tree produced by DmSqlParser#at_raft.
	VisitAt_raft(ctx *At_raftContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_trigger_stmt.
	VisitCreate_trigger_stmt(ctx *Create_trigger_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#before_after.
	VisitBefore_after(ctx *Before_afterContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_del_ins_upd_list.
	VisitTrig_del_ins_upd_list(ctx *Trig_del_ins_upd_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_del_ins_upd.
	VisitTrig_del_ins_upd(ctx *Trig_del_ins_updContext) interface{}

	// Visit a parse tree produced by DmSqlParser#update_of_option.
	VisitUpdate_of_option(ctx *Update_of_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#nowait.
	VisitNowait(ctx *NowaitContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_event_list.
	VisitTrig_event_list(ctx *Trig_event_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_event.
	VisitTrig_event(ctx *Trig_eventContext) interface{}

	// Visit a parse tree produced by DmSqlParser#event_object.
	VisitEvent_object(ctx *Event_objectContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_referencing_def_options.
	VisitTrig_referencing_def_options(ctx *Trig_referencing_def_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_referencing_def.
	VisitTrig_referencing_def(ctx *Trig_referencing_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_referencing_list.
	VisitTrig_referencing_list(ctx *Trig_referencing_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_referencing_old.
	VisitTrig_referencing_old(ctx *Trig_referencing_oldContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_referencing_new.
	VisitTrig_referencing_new(ctx *Trig_referencing_newContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_for_each_option.
	VisitTrig_for_each_option(ctx *Trig_for_each_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_timer_rate.
	VisitTrig_timer_rate(ctx *Trig_timer_rateContext) interface{}

	// Visit a parse tree produced by DmSqlParser#exec_ep_seqno.
	VisitExec_ep_seqno(ctx *Exec_ep_seqnoContext) interface{}

	// Visit a parse tree produced by DmSqlParser#rate_over_day.
	VisitRate_over_day(ctx *Rate_over_dayContext) interface{}

	// Visit a parse tree produced by DmSqlParser#month_rate.
	VisitMonth_rate(ctx *Month_rateContext) interface{}

	// Visit a parse tree produced by DmSqlParser#day_in_month.
	VisitDay_in_month(ctx *Day_in_monthContext) interface{}

	// Visit a parse tree produced by DmSqlParser#day_in_month_week.
	VisitDay_in_month_week(ctx *Day_in_month_weekContext) interface{}

	// Visit a parse tree produced by DmSqlParser#week_rate.
	VisitWeek_rate(ctx *Week_rateContext) interface{}

	// Visit a parse tree produced by DmSqlParser#day_of_week_list.
	VisitDay_of_week_list(ctx *Day_of_week_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#day_rate.
	VisitDay_rate(ctx *Day_rateContext) interface{}

	// Visit a parse tree produced by DmSqlParser#rate_in_day.
	VisitRate_in_day(ctx *Rate_in_dayContext) interface{}

	// Visit a parse tree produced by DmSqlParser#once_in_day.
	VisitOnce_in_day(ctx *Once_in_dayContext) interface{}

	// Visit a parse tree produced by DmSqlParser#times_in_day.
	VisitTimes_in_day(ctx *Times_in_dayContext) interface{}

	// Visit a parse tree produced by DmSqlParser#duaring_time.
	VisitDuaring_time(ctx *Duaring_timeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#duaring_date.
	VisitDuaring_date(ctx *Duaring_dateContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_when_option.
	VisitTrig_when_option(ctx *Trig_when_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_when_condition.
	VisitTrig_when_condition(ctx *Trig_when_conditionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#repeat_interval_stmt.
	VisitRepeat_interval_stmt(ctx *Repeat_interval_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#max_run_duration.
	VisitMax_run_duration(ctx *Max_run_durationContext) interface{}

	// Visit a parse tree produced by DmSqlParser#repeat_interval.
	VisitRepeat_interval(ctx *Repeat_intervalContext) interface{}

	// Visit a parse tree produced by DmSqlParser#frequency_clause.
	VisitFrequency_clause(ctx *Frequency_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#frequency_define.
	VisitFrequency_define(ctx *Frequency_defineContext) interface{}

	// Visit a parse tree produced by DmSqlParser#predefined_frequency.
	VisitPredefined_frequency(ctx *Predefined_frequencyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#interval_clause_list.
	VisitInterval_clause_list(ctx *Interval_clause_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#interval_clause_single.
	VisitInterval_clause_single(ctx *Interval_clause_singleContext) interface{}

	// Visit a parse tree produced by DmSqlParser#interval_clause.
	VisitInterval_clause(ctx *Interval_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#intervalnum.
	VisitIntervalnum(ctx *IntervalnumContext) interface{}

	// Visit a parse tree produced by DmSqlParser#bymonth_clause.
	VisitBymonth_clause(ctx *Bymonth_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#monthlist.
	VisitMonthlist(ctx *MonthlistContext) interface{}

	// Visit a parse tree produced by DmSqlParser#month.
	VisitMonth(ctx *MonthContext) interface{}

	// Visit a parse tree produced by DmSqlParser#numeric_month.
	VisitNumeric_month(ctx *Numeric_monthContext) interface{}

	// Visit a parse tree produced by DmSqlParser#char_month.
	VisitChar_month(ctx *Char_monthContext) interface{}

	// Visit a parse tree produced by DmSqlParser#byweekno_clause.
	VisitByweekno_clause(ctx *Byweekno_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#weekno_list.
	VisitWeekno_list(ctx *Weekno_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#weekno.
	VisitWeekno(ctx *WeeknoContext) interface{}

	// Visit a parse tree produced by DmSqlParser#byyearday_clause.
	VisitByyearday_clause(ctx *Byyearday_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#yearday_list.
	VisitYearday_list(ctx *Yearday_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#yearday.
	VisitYearday(ctx *YeardayContext) interface{}

	// Visit a parse tree produced by DmSqlParser#bymonthday_clause.
	VisitBymonthday_clause(ctx *Bymonthday_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#monthday_list.
	VisitMonthday_list(ctx *Monthday_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#monthday.
	VisitMonthday(ctx *MonthdayContext) interface{}

	// Visit a parse tree produced by DmSqlParser#byday_clause.
	VisitByday_clause(ctx *Byday_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#byday_list.
	VisitByday_list(ctx *Byday_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#byday.
	VisitByday(ctx *BydayContext) interface{}

	// Visit a parse tree produced by DmSqlParser#weekdaynum_options.
	VisitWeekdaynum_options(ctx *Weekdaynum_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#weekdaynum.
	VisitWeekdaynum(ctx *WeekdaynumContext) interface{}

	// Visit a parse tree produced by DmSqlParser#day.
	VisitDay(ctx *DayContext) interface{}

	// Visit a parse tree produced by DmSqlParser#byhour_clause.
	VisitByhour_clause(ctx *Byhour_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hour_list.
	VisitHour_list(ctx *Hour_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hour.
	VisitHour(ctx *HourContext) interface{}

	// Visit a parse tree produced by DmSqlParser#byminute_clause.
	VisitByminute_clause(ctx *Byminute_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#minute_list.
	VisitMinute_list(ctx *Minute_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#minute.
	VisitMinute(ctx *MinuteContext) interface{}

	// Visit a parse tree produced by DmSqlParser#bysecond_clause.
	VisitBysecond_clause(ctx *Bysecond_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#second_list.
	VisitSecond_list(ctx *Second_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#second.
	VisitSecond(ctx *SecondContext) interface{}

	// Visit a parse tree produced by DmSqlParser#query_rewrite.
	VisitQuery_rewrite(ctx *Query_rewriteContext) interface{}

	// Visit a parse tree produced by DmSqlParser#build_clause.
	VisitBuild_clause(ctx *Build_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mv_refresh_option.
	VisitMv_refresh_option(ctx *Mv_refresh_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mv_refresh_option_list.
	VisitMv_refresh_option_list(ctx *Mv_refresh_option_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mv_refresh_clause.
	VisitMv_refresh_clause(ctx *Mv_refresh_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mv_log_purge_syn_asyn_clause.
	VisitMv_log_purge_syn_asyn_clause(ctx *Mv_log_purge_syn_asyn_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mv_log_purge_clause.
	VisitMv_log_purge_clause(ctx *Mv_log_purge_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mv_log_with_syntax_item.
	VisitMv_log_with_syntax_item(ctx *Mv_log_with_syntax_itemContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mv_log_with_syntax_item_list.
	VisitMv_log_with_syntax_item_list(ctx *Mv_log_with_syntax_item_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mv_log_including_new_values.
	VisitMv_log_including_new_values(ctx *Mv_log_including_new_valuesContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mv_log_with_clause_null.
	VisitMv_log_with_clause_null(ctx *Mv_log_with_clause_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_materialized_view_log_stmt.
	VisitCreate_materialized_view_log_stmt(ctx *Create_materialized_view_log_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#prebuilt_table_clause_null.
	VisitPrebuilt_table_clause_null(ctx *Prebuilt_table_clause_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_materialized_view_stmt.
	VisitCreate_materialized_view_stmt(ctx *Create_materialized_view_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_view_stmt.
	VisitCreate_view_stmt(ctx *Create_view_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_view_stmt_body.
	VisitCreate_view_stmt_body(ctx *Create_view_stmt_bodyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_list3_options.
	VisitColumn_list3_options(ctx *Column_list3_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_list3.
	VisitColumn_list3(ctx *Column_list3Context) interface{}

	// Visit a parse tree produced by DmSqlParser#view_column_constraint_def.
	VisitView_column_constraint_def(ctx *View_column_constraint_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#view_column_constraints.
	VisitView_column_constraints(ctx *View_column_constraintsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#view_column_constraint.
	VisitView_column_constraint(ctx *View_column_constraintContext) interface{}

	// Visit a parse tree produced by DmSqlParser#view_column_constraint_action.
	VisitView_column_constraint_action(ctx *View_column_constraint_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#view_constraint_def.
	VisitView_constraint_def(ctx *View_constraint_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_schemabinding.
	VisitWith_schemabinding(ctx *With_schemabindingContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_list_options.
	VisitColumn_list_options(ctx *Column_list_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_check_or_readonly_option.
	VisitWith_check_or_readonly_option(ctx *With_check_or_readonly_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#check_level_option.
	VisitCheck_level_option(ctx *Check_level_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#decl_cursor.
	VisitDecl_cursor(ctx *Decl_cursorContext) interface{}

	// Visit a parse tree produced by DmSqlParser#delete_stmt.
	VisitDelete_stmt(ctx *Delete_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#delete_stmt_body.
	VisitDelete_stmt_body(ctx *Delete_stmt_bodyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#delete_multi_tv_option.
	VisitDelete_multi_tv_option(ctx *Delete_multi_tv_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#check_limit_option.
	VisitCheck_limit_option(ctx *Check_limit_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#where_current_option.
	VisitWhere_current_option(ctx *Where_current_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#where_clause.
	VisitWhere_clause(ctx *Where_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#start_with_clause_null.
	VisitStart_with_clause_null(ctx *Start_with_clause_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#connect_by_item.
	VisitConnect_by_item(ctx *Connect_by_itemContext) interface{}

	// Visit a parse tree produced by DmSqlParser#connect_by_clause.
	VisitConnect_by_clause(ctx *Connect_by_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hierarchical_query_clause.
	VisitHierarchical_query_clause(ctx *Hierarchical_query_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#nocycle_flag.
	VisitNocycle_flag(ctx *Nocycle_flagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#search_condition.
	VisitSearch_condition(ctx *Search_conditionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#disconnect_stmt.
	VisitDisconnect_stmt(ctx *Disconnect_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#disconnect_option.
	VisitDisconnect_option(ctx *Disconnect_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#drop_stmt.
	VisitDrop_stmt(ctx *Drop_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#drop_stmt_body_1.
	VisitDrop_stmt_body_1(ctx *Drop_stmt_body_1Context) interface{}

	// Visit a parse tree produced by DmSqlParser#drop_stmt_2.
	VisitDrop_stmt_2(ctx *Drop_stmt_2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#drop_id_db_object.
	VisitDrop_id_db_object(ctx *Drop_id_db_objectContext) interface{}

	// Visit a parse tree produced by DmSqlParser#id_db_object.
	VisitId_db_object(ctx *Id_db_objectContext) interface{}

	// Visit a parse tree produced by DmSqlParser#drop_db_object.
	VisitDrop_db_object(ctx *Drop_db_objectContext) interface{}

	// Visit a parse tree produced by DmSqlParser#exist.
	VisitExist(ctx *ExistContext) interface{}

	// Visit a parse tree produced by DmSqlParser#not_exist.
	VisitNot_exist(ctx *Not_existContext) interface{}

	// Visit a parse tree produced by DmSqlParser#db_object.
	VisitDb_object(ctx *Db_objectContext) interface{}

	// Visit a parse tree produced by DmSqlParser#is_detach.
	VisitIs_detach(ctx *Is_detachContext) interface{}

	// Visit a parse tree produced by DmSqlParser#purge_option.
	VisitPurge_option(ctx *Purge_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_database_stmt.
	VisitAlter_database_stmt(ctx *Alter_database_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_system_action.
	VisitAlter_system_action(ctx *Alter_system_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_database_action.
	VisitAlter_database_action(ctx *Alter_database_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#force.
	VisitForce(ctx *ForceContext) interface{}

	// Visit a parse tree produced by DmSqlParser#tablespace_name.
	VisitTablespace_name(ctx *Tablespace_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#raft_name.
	VisitRaft_name(ctx *Raft_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#fetch_into.
	VisitFetch_into(ctx *Fetch_intoContext) interface{}

	// Visit a parse tree produced by DmSqlParser#bulk_or_single_into.
	VisitBulk_or_single_into(ctx *Bulk_or_single_intoContext) interface{}

	// Visit a parse tree produced by DmSqlParser#fetch_stmt.
	VisitFetch_stmt(ctx *Fetch_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#fetch_statement.
	VisitFetch_statement(ctx *Fetch_statementContext) interface{}

	// Visit a parse tree produced by DmSqlParser#fetch_tail.
	VisitFetch_tail(ctx *Fetch_tailContext) interface{}

	// Visit a parse tree produced by DmSqlParser#fetch_limit_option.
	VisitFetch_limit_option(ctx *Fetch_limit_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#fetch_option.
	VisitFetch_option(ctx *Fetch_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#fetch_direct_option.
	VisitFetch_direct_option(ctx *Fetch_direct_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#log_errors_into.
	VisitLog_errors_into(ctx *Log_errors_intoContext) interface{}

	// Visit a parse tree produced by DmSqlParser#log_errors_expression.
	VisitLog_errors_expression(ctx *Log_errors_expressionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#log_errors_unlimited.
	VisitLog_errors_unlimited(ctx *Log_errors_unlimitedContext) interface{}

	// Visit a parse tree produced by DmSqlParser#log_errors.
	VisitLog_errors(ctx *Log_errorsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#insert_stmt.
	VisitInsert_stmt(ctx *Insert_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#insert_stmt_body.
	VisitInsert_stmt_body(ctx *Insert_stmt_bodyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_column_list_options.
	VisitFull_column_list_options(ctx *Full_column_list_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ins_value_options.
	VisitIns_value_options(ctx *Ins_value_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#insert_into_single.
	VisitInsert_into_single(ctx *Insert_into_singleContext) interface{}

	// Visit a parse tree produced by DmSqlParser#multi_insert_into_list.
	VisitMulti_insert_into_list(ctx *Multi_insert_into_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#multi_insert_tag.
	VisitMulti_insert_tag(ctx *Multi_insert_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#insert_into_single_condition.
	VisitInsert_into_single_condition(ctx *Insert_into_single_conditionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#multi_insert_into_condition_list.
	VisitMulti_insert_into_condition_list(ctx *Multi_insert_into_condition_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#multi_insert_into_else.
	VisitMulti_insert_into_else(ctx *Multi_insert_into_elseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#multi_insert_stmt_body.
	VisitMulti_insert_stmt_body(ctx *Multi_insert_stmt_bodyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#insert_tail.
	VisitInsert_tail(ctx *Insert_tailContext) interface{}

	// Visit a parse tree produced by DmSqlParser#insert_action.
	VisitInsert_action(ctx *Insert_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#record_var_values.
	VisitRecord_var_values(ctx *Record_var_valuesContext) interface{}

	// Visit a parse tree produced by DmSqlParser#record_var_value.
	VisitRecord_var_value(ctx *Record_var_valueContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ins_value.
	VisitIns_value(ctx *Ins_valueContext) interface{}

	// Visit a parse tree produced by DmSqlParser#open_stmt.
	VisitOpen_stmt(ctx *Open_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#open_statement.
	VisitOpen_statement(ctx *Open_statementContext) interface{}

	// Visit a parse tree produced by DmSqlParser#open_tail.
	VisitOpen_tail(ctx *Open_tailContext) interface{}

	// Visit a parse tree produced by DmSqlParser#return_stmt.
	VisitReturn_stmt(ctx *Return_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#raise_stmt.
	VisitRaise_stmt(ctx *Raise_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#rollback_stmt.
	VisitRollback_stmt(ctx *Rollback_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#to_savepoint.
	VisitTo_savepoint(ctx *To_savepointContext) interface{}

	// Visit a parse tree produced by DmSqlParser#savepoint_stmt.
	VisitSavepoint_stmt(ctx *Savepoint_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#select_stmt.
	VisitSelect_stmt(ctx *Select_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#all_distinct_option.
	VisitAll_distinct_option(ctx *All_distinct_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#all_distinct_option_2.
	VisitAll_distinct_option_2(ctx *All_distinct_option_2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#corresponding_clause.
	VisitCorresponding_clause(ctx *Corresponding_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#top_option.
	VisitTop_option(ctx *Top_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#limit_option.
	VisitLimit_option(ctx *Limit_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#limit_clause.
	VisitLimit_clause(ctx *Limit_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#limit_not_empty.
	VisitLimit_not_empty(ctx *Limit_not_emptyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#row_limiting_clause.
	VisitRow_limiting_clause(ctx *Row_limiting_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#row_num_desc.
	VisitRow_num_desc(ctx *Row_num_descContext) interface{}

	// Visit a parse tree produced by DmSqlParser#first_next_desc.
	VisitFirst_next_desc(ctx *First_next_descContext) interface{}

	// Visit a parse tree produced by DmSqlParser#select_item_list.
	VisitSelect_item_list(ctx *Select_item_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#select_item.
	VisitSelect_item(ctx *Select_itemContext) interface{}

	// Visit a parse tree produced by DmSqlParser#as_alias.
	VisitAs_alias(ctx *As_aliasContext) interface{}

	// Visit a parse tree produced by DmSqlParser#select_tail.
	VisitSelect_tail(ctx *Select_tailContext) interface{}

	// Visit a parse tree produced by DmSqlParser#from_clause.
	VisitFrom_clause(ctx *From_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#from_tv_list.
	VisitFrom_tv_list(ctx *From_tv_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#from_tv.
	VisitFrom_tv(ctx *From_tvContext) interface{}

	// Visit a parse tree produced by DmSqlParser#joined_table.
	VisitJoined_table(ctx *Joined_tableContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trxid.
	VisitTrxid(ctx *TrxidContext) interface{}

	// Visit a parse tree produced by DmSqlParser#flashback_query_low.
	VisitFlashback_query_low(ctx *Flashback_query_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trxid_option.
	VisitTrxid_option(ctx *Trxid_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#range_from_to.
	VisitRange_from_to(ctx *Range_from_toContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sample_exp.
	VisitSample_exp(ctx *Sample_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pivot_sfun.
	VisitPivot_sfun(ctx *Pivot_sfunContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pivot_sfun_lst.
	VisitPivot_sfun_lst(ctx *Pivot_sfun_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pivot_for_clause.
	VisitPivot_for_clause(ctx *Pivot_for_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pivot_in_clause1_expr.
	VisitPivot_in_clause1_expr(ctx *Pivot_in_clause1_exprContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pivot_in_clause_low_1.
	VisitPivot_in_clause_low_1(ctx *Pivot_in_clause_low_1Context) interface{}

	// Visit a parse tree produced by DmSqlParser#pivot_in_clause_low_2.
	VisitPivot_in_clause_low_2(ctx *Pivot_in_clause_low_2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#pivot_in_clause_low.
	VisitPivot_in_clause_low(ctx *Pivot_in_clause_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pivot_xml.
	VisitPivot_xml(ctx *Pivot_xmlContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pivot_clause_low.
	VisitPivot_clause_low(ctx *Pivot_clause_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#unpivot_val_col_lst.
	VisitUnpivot_val_col_lst(ctx *Unpivot_val_col_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#include_clause.
	VisitInclude_clause(ctx *Include_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#unpivot_in_clause_expr.
	VisitUnpivot_in_clause_expr(ctx *Unpivot_in_clause_exprContext) interface{}

	// Visit a parse tree produced by DmSqlParser#unpivot_in_clause_low.
	VisitUnpivot_in_clause_low(ctx *Unpivot_in_clause_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#unpivot_clause_low.
	VisitUnpivot_clause_low(ctx *Unpivot_clause_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sample_clause_low.
	VisitSample_clause_low(ctx *Sample_clause_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#normal_tv_name.
	VisitNormal_tv_name(ctx *Normal_tv_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#normal_tv_tail.
	VisitNormal_tv_tail(ctx *Normal_tv_tailContext) interface{}

	// Visit a parse tree produced by DmSqlParser#normal_tv_tail_low.
	VisitNormal_tv_tail_low(ctx *Normal_tv_tail_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#normal_alias.
	VisitNormal_alias(ctx *Normal_aliasContext) interface{}

	// Visit a parse tree produced by DmSqlParser#normal_tv_tail_low2.
	VisitNormal_tv_tail_low2(ctx *Normal_tv_tail_low2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#normal_tv_tail_low3.
	VisitNormal_tv_tail_low3(ctx *Normal_tv_tail_low3Context) interface{}

	// Visit a parse tree produced by DmSqlParser#normal_tv_derived_table_options.
	VisitNormal_tv_derived_table_options(ctx *Normal_tv_derived_table_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#normal_tv_derived_table_low.
	VisitNormal_tv_derived_table_low(ctx *Normal_tv_derived_table_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#normal_tv_derived_table.
	VisitNormal_tv_derived_table(ctx *Normal_tv_derived_tableContext) interface{}

	// Visit a parse tree produced by DmSqlParser#select_with_paran_with_alias.
	VisitSelect_with_paran_with_alias(ctx *Select_with_paran_with_aliasContext) interface{}

	// Visit a parse tree produced by DmSqlParser#from_table_exp.
	VisitFrom_table_exp(ctx *From_table_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#from_table_select_with_paran.
	VisitFrom_table_select_with_paran(ctx *From_table_select_with_paranContext) interface{}

	// Visit a parse tree produced by DmSqlParser#normal_tv.
	VisitNormal_tv(ctx *Normal_tvContext) interface{}

	// Visit a parse tree produced by DmSqlParser#xml_passing.
	VisitXml_passing(ctx *Xml_passingContext) interface{}

	// Visit a parse tree produced by DmSqlParser#xmlcoldef_lst_options.
	VisitXmlcoldef_lst_options(ctx *Xmlcoldef_lst_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#xmlcoldef_lst.
	VisitXmlcoldef_lst(ctx *Xmlcoldef_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#xmlcoldef.
	VisitXmlcoldef(ctx *XmlcoldefContext) interface{}

	// Visit a parse tree produced by DmSqlParser#on_error.
	VisitOn_error(ctx *On_errorContext) interface{}

	// Visit a parse tree produced by DmSqlParser#jsoncol_lst.
	VisitJsoncol_lst(ctx *Jsoncol_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#jsoncoldef_lst.
	VisitJsoncoldef_lst(ctx *Jsoncoldef_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#jsoncoldef.
	VisitJsoncoldef(ctx *JsoncoldefContext) interface{}

	// Visit a parse tree produced by DmSqlParser#json_exists_col.
	VisitJson_exists_col(ctx *Json_exists_colContext) interface{}

	// Visit a parse tree produced by DmSqlParser#json_qurey_col.
	VisitJson_qurey_col(ctx *Json_qurey_colContext) interface{}

	// Visit a parse tree produced by DmSqlParser#json_value_col.
	VisitJson_value_col(ctx *Json_value_colContext) interface{}

	// Visit a parse tree produced by DmSqlParser#json_nested_col.
	VisitJson_nested_col(ctx *Json_nested_colContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ordinality_col.
	VisitOrdinality_col(ctx *Ordinality_colContext) interface{}

	// Visit a parse tree produced by DmSqlParser#joined_table_element.
	VisitJoined_table_element(ctx *Joined_table_elementContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cross_outer_apply_clause.
	VisitCross_outer_apply_clause(ctx *Cross_outer_apply_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cross_outer_apply_join.
	VisitCross_outer_apply_join(ctx *Cross_outer_apply_joinContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cross_join.
	VisitCross_join(ctx *Cross_joinContext) interface{}

	// Visit a parse tree produced by DmSqlParser#partition_out_join.
	VisitPartition_out_join(ctx *Partition_out_joinContext) interface{}

	// Visit a parse tree produced by DmSqlParser#qualified_join.
	VisitQualified_join(ctx *Qualified_joinContext) interface{}

	// Visit a parse tree produced by DmSqlParser#qualified_joinspec.
	VisitQualified_joinspec(ctx *Qualified_joinspecContext) interface{}

	// Visit a parse tree produced by DmSqlParser#named_columns_join.
	VisitNamed_columns_join(ctx *Named_columns_joinContext) interface{}

	// Visit a parse tree produced by DmSqlParser#join_hint.
	VisitJoin_hint(ctx *Join_hintContext) interface{}

	// Visit a parse tree produced by DmSqlParser#join_type.
	VisitJoin_type(ctx *Join_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#outer_join_type.
	VisitOuter_join_type(ctx *Outer_join_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#join_condition.
	VisitJoin_condition(ctx *Join_conditionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#group_by_clause.
	VisitGroup_by_clause(ctx *Group_by_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#group_item.
	VisitGroup_item(ctx *Group_itemContext) interface{}

	// Visit a parse tree produced by DmSqlParser#exp_rollup_cube_item2.
	VisitExp_rollup_cube_item2(ctx *Exp_rollup_cube_item2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#exp_rollup_cube_item.
	VisitExp_rollup_cube_item(ctx *Exp_rollup_cube_itemContext) interface{}

	// Visit a parse tree produced by DmSqlParser#grouping_set_items.
	VisitGrouping_set_items(ctx *Grouping_set_itemsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#grouping_set_item.
	VisitGrouping_set_item(ctx *Grouping_set_itemContext) interface{}

	// Visit a parse tree produced by DmSqlParser#having_clause.
	VisitHaving_clause(ctx *Having_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#without_into_select.
	VisitWithout_into_select(ctx *Without_into_selectContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sel_clause_app.
	VisitSel_clause_app(ctx *Sel_clause_appContext) interface{}

	// Visit a parse tree produced by DmSqlParser#select_clause.
	VisitSelect_clause(ctx *Select_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#simple_select.
	VisitSimple_select(ctx *Simple_selectContext) interface{}

	// Visit a parse tree produced by DmSqlParser#select_with_paran.
	VisitSelect_with_paran(ctx *Select_with_paranContext) interface{}

	// Visit a parse tree produced by DmSqlParser#query_exp.
	VisitQuery_exp(ctx *Query_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#for_xml_path.
	VisitFor_xml_path(ctx *For_xml_pathContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_tag.
	VisitWith_tag(ctx *With_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_option.
	VisitWith_option(ctx *With_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_clause.
	VisitWith_clause(ctx *With_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_function_list.
	VisitWith_function_list(ctx *With_function_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#func_def_inner.
	VisitFunc_def_inner(ctx *Func_def_innerContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_function_list_element.
	VisitWith_function_list_element(ctx *With_function_list_elementContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_view_list.
	VisitWith_view_list(ctx *With_view_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#depth_type_option.
	VisitDepth_type_option(ctx *Depth_type_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#search_clause.
	VisitSearch_clause(ctx *Search_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cycle_clause.
	VisitCycle_clause(ctx *Cycle_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_view_list_element.
	VisitWith_view_list_element(ctx *With_view_list_elementContext) interface{}

	// Visit a parse tree produced by DmSqlParser#assignment_obj_list.
	VisitAssignment_obj_list(ctx *Assignment_obj_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#assignment_obj.
	VisitAssignment_obj(ctx *Assignment_objContext) interface{}

	// Visit a parse tree produced by DmSqlParser#order_by_options.
	VisitOrder_by_options(ctx *Order_by_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#order_by.
	VisitOrder_by(ctx *Order_byContext) interface{}

	// Visit a parse tree produced by DmSqlParser#asc_desc_option.
	VisitAsc_desc_option(ctx *Asc_desc_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#nulls_last_option.
	VisitNulls_last_option(ctx *Nulls_last_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#collate_option.
	VisitCollate_option(ctx *Collate_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#order_by_list.
	VisitOrder_by_list(ctx *Order_by_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#order_by_item.
	VisitOrder_by_item(ctx *Order_by_itemContext) interface{}

	// Visit a parse tree produced by DmSqlParser#for_update_options.
	VisitFor_update_options(ctx *For_update_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#for_update.
	VisitFor_update(ctx *For_updateContext) interface{}

	// Visit a parse tree produced by DmSqlParser#set_session_stmt.
	VisitSet_session_stmt(ctx *Set_session_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#set_trans_stmt.
	VisitSet_trans_stmt(ctx *Set_trans_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trans_mode_lstl.
	VisitTrans_mode_lstl(ctx *Trans_mode_lstlContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trans_mode_lst.
	VisitTrans_mode_lst(ctx *Trans_mode_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trans_mode.
	VisitTrans_mode(ctx *Trans_modeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#time_zone_exp_new.
	VisitTime_zone_exp_new(ctx *Time_zone_exp_newContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trans_rw_option.
	VisitTrans_rw_option(ctx *Trans_rw_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trans_level_option.
	VisitTrans_level_option(ctx *Trans_level_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#lock_table_stmt.
	VisitLock_table_stmt(ctx *Lock_table_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#lock_mode_option.
	VisitLock_mode_option(ctx *Lock_mode_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#set_identins_stmt.
	VisitSet_identins_stmt(ctx *Set_identins_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#set_identins_option.
	VisitSet_identins_option(ctx *Set_identins_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trunc_table_stmt.
	VisitTrunc_table_stmt(ctx *Trunc_table_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#update_stmt.
	VisitUpdate_stmt(ctx *Update_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#update_stmt_body.
	VisitUpdate_stmt_body(ctx *Update_stmt_bodyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#update_tv_list.
	VisitUpdate_tv_list(ctx *Update_tv_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#return_item.
	VisitReturn_item(ctx *Return_itemContext) interface{}

	// Visit a parse tree produced by DmSqlParser#return_item_list.
	VisitReturn_item_list(ctx *Return_item_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#return_option.
	VisitReturn_option(ctx *Return_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#return_into_obj.
	VisitReturn_into_obj(ctx *Return_into_objContext) interface{}

	// Visit a parse tree produced by DmSqlParser#collect_into_rset.
	VisitCollect_into_rset(ctx *Collect_into_rsetContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alias_option.
	VisitAlias_option(ctx *Alias_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#set_value_list.
	VisitSet_value_list(ctx *Set_value_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#set_value_node.
	VisitSet_value_node(ctx *Set_value_nodeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#set_col_list.
	VisitSet_col_list(ctx *Set_col_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#update_from_clause.
	VisitUpdate_from_clause(ctx *Update_from_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#merge_into_stmt.
	VisitMerge_into_stmt(ctx *Merge_into_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#merge_into_stmt_body.
	VisitMerge_into_stmt_body(ctx *Merge_into_stmt_bodyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#merge_into_sub_clause.
	VisitMerge_into_sub_clause(ctx *Merge_into_sub_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#merge_update_clause.
	VisitMerge_update_clause(ctx *Merge_update_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#merge_insert_clause.
	VisitMerge_insert_clause(ctx *Merge_insert_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_profile_stmt.
	VisitCreate_profile_stmt(ctx *Create_profile_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_profile_stmt.
	VisitAlter_profile_stmt(ctx *Alter_profile_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_user_stmt.
	VisitCreate_user_stmt(ctx *Create_user_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#default_ts_name.
	VisitDefault_ts_name(ctx *Default_ts_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#default_ts_name_lst.
	VisitDefault_ts_name_lst(ctx *Default_ts_name_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#default_ts_name_node.
	VisitDefault_ts_name_node(ctx *Default_ts_name_nodeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#default_idx_ts_name.
	VisitDefault_idx_ts_name(ctx *Default_idx_ts_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#default_ts_name_low.
	VisitDefault_ts_name_low(ctx *Default_ts_name_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#temp_ts_name.
	VisitTemp_ts_name(ctx *Temp_ts_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#default_ts_group_name_low.
	VisitDefault_ts_group_name_low(ctx *Default_ts_group_name_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#on_schema.
	VisitOn_schema(ctx *On_schemaContext) interface{}

	// Visit a parse tree produced by DmSqlParser#replace_old_pwd.
	VisitReplace_old_pwd(ctx *Replace_old_pwdContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_user_stmt.
	VisitAlter_user_stmt(ctx *Alter_user_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#user_encrypt_options.
	VisitUser_encrypt_options(ctx *User_encrypt_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#user_encrypt_option.
	VisitUser_encrypt_option(ctx *User_encrypt_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#authent_type_options.
	VisitAuthent_type_options(ctx *Authent_type_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#hash_cipher_option.
	VisitHash_cipher_option(ctx *Hash_cipher_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#authent_type.
	VisitAuthent_type(ctx *Authent_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#force_format.
	VisitForce_format(ctx *Force_formatContext) interface{}

	// Visit a parse tree produced by DmSqlParser#as.
	VisitAs(ctx *AsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pwd_policy.
	VisitPwd_policy(ctx *Pwd_policyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#account_lock.
	VisitAccount_lock(ctx *Account_lockContext) interface{}

	// Visit a parse tree produced by DmSqlParser#read_only_flag.
	VisitRead_only_flag(ctx *Read_only_flagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#read_only_flag_not_null.
	VisitRead_only_flag_not_null(ctx *Read_only_flag_not_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#resource.
	VisitResource(ctx *ResourceContext) interface{}

	// Visit a parse tree produced by DmSqlParser#expire.
	VisitExpire(ctx *ExpireContext) interface{}

	// Visit a parse tree produced by DmSqlParser#resource_limit_options.
	VisitResource_limit_options(ctx *Resource_limit_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#resource_limit_list.
	VisitResource_limit_list(ctx *Resource_limit_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#resource_limit_list_with_comma.
	VisitResource_limit_list_with_comma(ctx *Resource_limit_list_with_commaContext) interface{}

	// Visit a parse tree produced by DmSqlParser#resource_limit_list_with_empty.
	VisitResource_limit_list_with_empty(ctx *Resource_limit_list_with_emptyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#resource_limit.
	VisitResource_limit(ctx *Resource_limitContext) interface{}

	// Visit a parse tree produced by DmSqlParser#resource_limit_value.
	VisitResource_limit_value(ctx *Resource_limit_valueContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_audit_rule_stmt.
	VisitCreate_audit_rule_stmt(ctx *Create_audit_rule_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#rule_name.
	VisitRule_name(ctx *Rule_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#audit_rule_action.
	VisitAudit_rule_action(ctx *Audit_rule_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#by_login_or_all.
	VisitBy_login_or_all(ctx *By_login_or_allContext) interface{}

	// Visit a parse tree produced by DmSqlParser#allow_ip_list.
	VisitAllow_ip_list(ctx *Allow_ip_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#not_allow_ip_list.
	VisitNot_allow_ip_list(ctx *Not_allow_ip_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ip_list.
	VisitIp_list(ctx *Ip_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#allow_dt_list.
	VisitAllow_dt_list(ctx *Allow_dt_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#not_allow_dt_list.
	VisitNot_allow_dt_list(ctx *Not_allow_dt_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dt_list.
	VisitDt_list(ctx *Dt_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dt.
	VisitDt(ctx *DtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#op_freq.
	VisitOp_freq(ctx *Op_freqContext) interface{}

	// Visit a parse tree produced by DmSqlParser#quota_val.
	VisitQuota_val(ctx *Quota_valContext) interface{}

	// Visit a parse tree produced by DmSqlParser#quota_ts_node.
	VisitQuota_ts_node(ctx *Quota_ts_nodeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#quota_ts_lst.
	VisitQuota_ts_lst(ctx *Quota_ts_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#quota_ts.
	VisitQuota_ts(ctx *Quota_tsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_role_stmt.
	VisitCreate_role_stmt(ctx *Create_role_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_dblink_stmt.
	VisitCreate_dblink_stmt(ctx *Create_dblink_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#db_type_str.
	VisitDb_type_str(ctx *Db_type_strContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dblink_option_lst_options.
	VisitDblink_option_lst_options(ctx *Dblink_option_lst_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dblink_option_lst.
	VisitDblink_option_lst(ctx *Dblink_option_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dblink_option.
	VisitDblink_option(ctx *Dblink_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_synonym_stmt.
	VisitCreate_synonym_stmt(ctx *Create_synonym_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_synonym_name.
	VisitFull_synonym_name(ctx *Full_synonym_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_obj_name.
	VisitFull_obj_name(ctx *Full_obj_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_domain_stmt.
	VisitCreate_domain_stmt(ctx *Create_domain_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#domain_default_option.
	VisitDomain_default_option(ctx *Domain_default_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#domain_constraints_option.
	VisitDomain_constraints_option(ctx *Domain_constraints_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#domain_constraints_def.
	VisitDomain_constraints_def(ctx *Domain_constraints_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#domain_constraints.
	VisitDomain_constraints(ctx *Domain_constraintsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#domain_constraint.
	VisitDomain_constraint(ctx *Domain_constraintContext) interface{}

	// Visit a parse tree produced by DmSqlParser#domain_constraint_name_def_options.
	VisitDomain_constraint_name_def_options(ctx *Domain_constraint_name_def_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#domain_constraint_name_def.
	VisitDomain_constraint_name_def(ctx *Domain_constraint_name_defContext) interface{}

	// Visit a parse tree produced by DmSqlParser#domain_constraint_name.
	VisitDomain_constraint_name(ctx *Domain_constraint_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_character_set_stmt.
	VisitCreate_character_set_stmt(ctx *Create_character_set_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#character_set_source.
	VisitCharacter_set_source(ctx *Character_set_sourceContext) interface{}

	// Visit a parse tree produced by DmSqlParser#existing_character_set_name.
	VisitExisting_character_set_name(ctx *Existing_character_set_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#character_set_name.
	VisitCharacter_set_name(ctx *Character_set_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#collate_clause_option.
	VisitCollate_clause_option(ctx *Collate_clause_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#collation_name.
	VisitCollation_name(ctx *Collation_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_collation_stmt.
	VisitCreate_collation_stmt(ctx *Create_collation_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#existing_collation_name.
	VisitExisting_collation_name(ctx *Existing_collation_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pad_option.
	VisitPad_option(ctx *Pad_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_sequence_stmt.
	VisitCreate_sequence_stmt(ctx *Create_sequence_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sequence_option_list_options.
	VisitSequence_option_list_options(ctx *Sequence_option_list_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sequence_option_list.
	VisitSequence_option_list(ctx *Sequence_option_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sequence_option.
	VisitSequence_option(ctx *Sequence_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sequence_name.
	VisitSequence_name(ctx *Sequence_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#increment_option.
	VisitIncrement_option(ctx *Increment_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#start_option.
	VisitStart_option(ctx *Start_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#current_option.
	VisitCurrent_option(ctx *Current_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#maxvalue_option.
	VisitMaxvalue_option(ctx *Maxvalue_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#minvalue_option.
	VisitMinvalue_option(ctx *Minvalue_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cycle_option.
	VisitCycle_option(ctx *Cycle_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cache_option.
	VisitCache_option(ctx *Cache_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#order_option.
	VisitOrder_option(ctx *Order_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#seq_local_option.
	VisitSeq_local_option(ctx *Seq_local_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#whenever_stmt_options.
	VisitWhenever_stmt_options(ctx *Whenever_stmt_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#whenever_stmt.
	VisitWhenever_stmt(ctx *Whenever_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#grant_stmt.
	VisitGrant_stmt(ctx *Grant_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#grant_tag.
	VisitGrant_tag(ctx *Grant_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#grant_stmt_body.
	VisitGrant_stmt_body(ctx *Grant_stmt_bodyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#revoke_stmt.
	VisitRevoke_stmt(ctx *Revoke_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#revoke_stmt_body.
	VisitRevoke_stmt_body(ctx *Revoke_stmt_bodyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#grant_privs.
	VisitGrant_privs(ctx *Grant_privsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#grant_priv_list.
	VisitGrant_priv_list(ctx *Grant_priv_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#grant_priv_off.
	VisitGrant_priv_off(ctx *Grant_priv_offContext) interface{}

	// Visit a parse tree produced by DmSqlParser#grant_priv.
	VisitGrant_priv(ctx *Grant_privContext) interface{}

	// Visit a parse tree produced by DmSqlParser#revoke_action.
	VisitRevoke_action(ctx *Revoke_actionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#db_priv_list.
	VisitDb_priv_list(ctx *Db_priv_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#db_priv.
	VisitDb_priv(ctx *Db_privContext) interface{}

	// Visit a parse tree produced by DmSqlParser#by_grantor.
	VisitBy_grantor(ctx *By_grantorContext) interface{}

	// Visit a parse tree produced by DmSqlParser#grantees.
	VisitGrantees(ctx *GranteesContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_schema_stmt.
	VisitCreate_schema_stmt(ctx *Create_schema_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#oprt_arg.
	VisitOprt_arg(ctx *Oprt_argContext) interface{}

	// Visit a parse tree produced by DmSqlParser#oprt_arg_lst.
	VisitOprt_arg_lst(ctx *Oprt_arg_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_operator_stmt.
	VisitCreate_operator_stmt(ctx *Create_operator_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#drop_operator_stmt.
	VisitDrop_operator_stmt(ctx *Drop_operator_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#grant_and_ddl.
	VisitGrant_and_ddl(ctx *Grant_and_ddlContext) interface{}

	// Visit a parse tree produced by DmSqlParser#top_exp.
	VisitTop_exp(ctx *Top_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#u_oprt.
	VisitU_oprt(ctx *U_oprtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#qualified_u_oprt.
	VisitQualified_u_oprt(ctx *Qualified_u_oprtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#exp_u_oprt.
	VisitExp_u_oprt(ctx *Exp_u_oprtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#raw_exp.
	VisitRaw_exp(ctx *Raw_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#exp.
	VisitExp(ctx *ExpContext) interface{}

	// Visit a parse tree produced by DmSqlParser#from_first_last_option.
	VisitFrom_first_last_option(ctx *From_first_last_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#afun_arg_lst.
	VisitAfun_arg_lst(ctx *Afun_arg_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#afun_arg_lst_low.
	VisitAfun_arg_lst_low(ctx *Afun_arg_lst_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#in_value_exp.
	VisitIn_value_exp(ctx *In_value_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#afun_partition_by.
	VisitAfun_partition_by(ctx *Afun_partition_byContext) interface{}

	// Visit a parse tree produced by DmSqlParser#afun_windowing.
	VisitAfun_windowing(ctx *Afun_windowingContext) interface{}

	// Visit a parse tree produced by DmSqlParser#afun_windowing_type.
	VisitAfun_windowing_type(ctx *Afun_windowing_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#afun_range_clause.
	VisitAfun_range_clause(ctx *Afun_range_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pexp.
	VisitPexp(ctx *PexpContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pexp_pfx.
	VisitPexp_pfx(ctx *Pexp_pfxContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pexp_cast.
	VisitPexp_cast(ctx *Pexp_castContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pexp_b.
	VisitPexp_b(ctx *Pexp_bContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pexp_a.
	VisitPexp_a(ctx *Pexp_aContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pexp_c.
	VisitPexp_c(ctx *Pexp_cContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pexp_c_insert.
	VisitPexp_c_insert(ctx *Pexp_c_insertContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pexp_d.
	VisitPexp_d(ctx *Pexp_dContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pexp_e.
	VisitPexp_e(ctx *Pexp_eContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pexp_pfx2.
	VisitPexp_pfx2(ctx *Pexp_pfx2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#member2.
	VisitMember2(ctx *Member2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#pexp_c2_insert.
	VisitPexp_c2_insert(ctx *Pexp_c2_insertContext) interface{}

	// Visit a parse tree produced by DmSqlParser#member_access2.
	VisitMember_access2(ctx *Member_access2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#invocation_expression2.
	VisitInvocation_expression2(ctx *Invocation_expression2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#member.
	VisitMember(ctx *MemberContext) interface{}

	// Visit a parse tree produced by DmSqlParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#member_access.
	VisitMember_access(ctx *Member_accessContext) interface{}

	// Visit a parse tree produced by DmSqlParser#invocation_expression.
	VisitInvocation_expression(ctx *Invocation_expressionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#invocation_expression_low.
	VisitInvocation_expression_low(ctx *Invocation_expression_lowContext) interface{}

	// Visit a parse tree produced by DmSqlParser#xmlagg_inv_expression.
	VisitXmlagg_inv_expression(ctx *Xmlagg_inv_expressionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#xmlfun_inv_expression.
	VisitXmlfun_inv_expression(ctx *Xmlfun_inv_expressionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#xmlele_name.
	VisitXmlele_name(ctx *Xmlele_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#xmlele_sub_lst.
	VisitXmlele_sub_lst(ctx *Xmlele_sub_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#xmlattr_lst.
	VisitXmlattr_lst(ctx *Xmlattr_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#xmlattr.
	VisitXmlattr(ctx *XmlattrContext) interface{}

	// Visit a parse tree produced by DmSqlParser#xmlval_lst.
	VisitXmlval_lst(ctx *Xmlval_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#keep_clause.
	VisitKeep_clause(ctx *Keep_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#within_clause.
	VisitWithin_clause(ctx *Within_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#typeof_expression.
	VisitTypeof_expression(ctx *Typeof_expressionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#new_obj_expression.
	VisitNew_obj_expression(ctx *New_obj_expressionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#new_arr_expression.
	VisitNew_arr_expression(ctx *New_arr_expressionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#array_creation_expression.
	VisitArray_creation_expression(ctx *Array_creation_expressionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#plsql_datatype_ex.
	VisitPlsql_datatype_ex(ctx *Plsql_datatype_exContext) interface{}

	// Visit a parse tree produced by DmSqlParser#new_array_type.
	VisitNew_array_type(ctx *New_array_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#opt_array_initializer.
	VisitOpt_array_initializer(ctx *Opt_array_initializerContext) interface{}

	// Visit a parse tree produced by DmSqlParser#array_initializer.
	VisitArray_initializer(ctx *Array_initializerContext) interface{}

	// Visit a parse tree produced by DmSqlParser#variable_initializer_list.
	VisitVariable_initializer_list(ctx *Variable_initializer_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#variable_initializer.
	VisitVariable_initializer(ctx *Variable_initializerContext) interface{}

	// Visit a parse tree produced by DmSqlParser#opt_comma.
	VisitOpt_comma(ctx *Opt_commaContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sizeof_expression.
	VisitSizeof_expression(ctx *Sizeof_expressionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#element_access.
	VisitElement_access(ctx *Element_accessContext) interface{}

	// Visit a parse tree produced by DmSqlParser#decode_case.
	VisitDecode_case(ctx *Decode_caseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#else_exp.
	VisitElse_exp(ctx *Else_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#boolean_case.
	VisitBoolean_case(ctx *Boolean_caseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#if_exp.
	VisitIf_exp(ctx *If_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#bool_when_list.
	VisitBool_when_list(ctx *Bool_when_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ops.
	VisitOps(ctx *OpsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#value_list.
	VisitValue_list(ctx *Value_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#in_value_list.
	VisitIn_value_list(ctx *In_value_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#value_list_set.
	VisitValue_list_set(ctx *Value_list_setContext) interface{}

	// Visit a parse tree produced by DmSqlParser#comma_list.
	VisitComma_list(ctx *Comma_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ins_value_list.
	VisitIns_value_list(ctx *Ins_value_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#null_value.
	VisitNull_value(ctx *Null_valueContext) interface{}

	// Visit a parse tree produced by DmSqlParser#id_and_rsvd_word_others.
	VisitId_and_rsvd_word_others(ctx *Id_and_rsvd_word_othersContext) interface{}

	// Visit a parse tree produced by DmSqlParser#id_and_rsvd_word.
	VisitId_and_rsvd_word(ctx *Id_and_rsvd_wordContext) interface{}

	// Visit a parse tree produced by DmSqlParser#stm_param.
	VisitStm_param(ctx *Stm_paramContext) interface{}

	// Visit a parse tree produced by DmSqlParser#stm_param_normal.
	VisitStm_param_normal(ctx *Stm_param_normalContext) interface{}

	// Visit a parse tree produced by DmSqlParser#stm_param_name.
	VisitStm_param_name(ctx *Stm_param_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#param_name_options.
	VisitParam_name_options(ctx *Param_name_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#contains_query_exp.
	VisitContains_query_exp(ctx *Contains_query_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#contains_query_exp_lst.
	VisitContains_query_exp_lst(ctx *Contains_query_exp_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#contains_exp.
	VisitContains_exp(ctx *Contains_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#strict_option.
	VisitStrict_option(ctx *Strict_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_unique_option.
	VisitWith_unique_option(ctx *With_unique_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#type_option.
	VisitType_option(ctx *Type_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#type_element.
	VisitType_element(ctx *Type_elementContext) interface{}

	// Visit a parse tree produced by DmSqlParser#type_element_list.
	VisitType_element_list(ctx *Type_element_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#bool_exp.
	VisitBool_exp(ctx *Bool_expContext) interface{}

	// Visit a parse tree produced by DmSqlParser#bool_exp_element.
	VisitBool_exp_element(ctx *Bool_exp_elementContext) interface{}

	// Visit a parse tree produced by DmSqlParser#query_any_options.
	VisitQuery_any_options(ctx *Query_any_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#global_var.
	VisitGlobal_var(ctx *Global_varContext) interface{}

	// Visit a parse tree produced by DmSqlParser#reserved_word.
	VisitReserved_word(ctx *Reserved_wordContext) interface{}

	// Visit a parse tree produced by DmSqlParser#new_none_reserved_word.
	VisitNew_none_reserved_word(ctx *New_none_reserved_wordContext) interface{}

	// Visit a parse tree produced by DmSqlParser#interval_nresvd_word.
	VisitInterval_nresvd_word(ctx *Interval_nresvd_wordContext) interface{}

	// Visit a parse tree produced by DmSqlParser#variable_resvd_word.
	VisitVariable_resvd_word(ctx *Variable_resvd_wordContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alias_resvd_word.
	VisitAlias_resvd_word(ctx *Alias_resvd_wordContext) interface{}

	// Visit a parse tree produced by DmSqlParser#schname_resvd_word.
	VisitSchname_resvd_word(ctx *Schname_resvd_wordContext) interface{}

	// Visit a parse tree produced by DmSqlParser#raw_id.
	VisitRaw_id(ctx *Raw_idContext) interface{}

	// Visit a parse tree produced by DmSqlParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by DmSqlParser#qualified_name.
	VisitQualified_name(ctx *Qualified_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#qualified_name2.
	VisitQualified_name2(ctx *Qualified_name2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#variable_name.
	VisitVariable_name(ctx *Variable_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#end_loop_label_null.
	VisitEnd_loop_label_null(ctx *End_loop_label_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#label_name_options.
	VisitLabel_name_options(ctx *Label_name_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#label_name.
	VisitLabel_name(ctx *Label_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#database_name.
	VisitDatabase_name(ctx *Database_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#backup_name.
	VisitBackup_name(ctx *Backup_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_proc_name.
	VisitFull_proc_name(ctx *Full_proc_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_proc_name2.
	VisitFull_proc_name2(ctx *Full_proc_name2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#full_fun_name.
	VisitFull_fun_name(ctx *Full_fun_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_table_name.
	VisitFull_table_name(ctx *Full_table_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_grp_name.
	VisitFull_grp_name(ctx *Full_grp_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_table_name2.
	VisitFull_table_name2(ctx *Full_table_name2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#full_partition_name.
	VisitFull_partition_name(ctx *Full_partition_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_schema_name.
	VisitFull_schema_name(ctx *Full_schema_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_name.
	VisitColumn_name(ctx *Column_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#constraint_name.
	VisitConstraint_name(ctx *Constraint_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_trigger_name.
	VisitFull_trigger_name(ctx *Full_trigger_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_trigger_name2.
	VisitFull_trigger_name2(ctx *Full_trigger_name2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#full_view_name.
	VisitFull_view_name(ctx *Full_view_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_view_name2.
	VisitFull_view_name2(ctx *Full_view_name2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#cursor_name.
	VisitCursor_name(ctx *Cursor_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trigger_name.
	VisitTrigger_name(ctx *Trigger_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#login_name.
	VisitLogin_name(ctx *Login_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#profile_name.
	VisitProfile_name(ctx *Profile_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#user_name.
	VisitUser_name(ctx *User_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#role_name.
	VisitRole_name(ctx *Role_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#user_role_name.
	VisitUser_role_name(ctx *User_role_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#role_name_list.
	VisitRole_name_list(ctx *Role_name_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_func_name.
	VisitFull_func_name(ctx *Full_func_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#param_name.
	VisitParam_name(ctx *Param_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#index_name.
	VisitIndex_name(ctx *Index_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#index_name2.
	VisitIndex_name2(ctx *Index_name2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_old_name.
	VisitTrig_old_name(ctx *Trig_old_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#trig_new_name.
	VisitTrig_new_name(ctx *Trig_new_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_tv_name.
	VisitFull_tv_name(ctx *Full_tv_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#full_object_name.
	VisitFull_object_name(ctx *Full_object_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#orient_option.
	VisitOrient_option(ctx *Orient_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#datepart.
	VisitDatepart(ctx *DatepartContext) interface{}

	// Visit a parse tree produced by DmSqlParser#datepart_op.
	VisitDatepart_op(ctx *Datepart_opContext) interface{}

	// Visit a parse tree produced by DmSqlParser#datead_fun.
	VisitDatead_fun(ctx *Datead_funContext) interface{}

	// Visit a parse tree produced by DmSqlParser#returning.
	VisitReturning(ctx *ReturningContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pretty.
	VisitPretty(ctx *PrettyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#wrapper_flag.
	VisitWrapper_flag(ctx *Wrapper_flagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#array_wrapper.
	VisitArray_wrapper(ctx *Array_wrapperContext) interface{}

	// Visit a parse tree produced by DmSqlParser#json_tail_on_empty.
	VisitJson_tail_on_empty(ctx *Json_tail_on_emptyContext) interface{}

	// Visit a parse tree produced by DmSqlParser#empty_handle.
	VisitEmpty_handle(ctx *Empty_handleContext) interface{}

	// Visit a parse tree produced by DmSqlParser#json_tail_on_error_null.
	VisitJson_tail_on_error_null(ctx *Json_tail_on_error_nullContext) interface{}

	// Visit a parse tree produced by DmSqlParser#error_handle.
	VisitError_handle(ctx *Error_handleContext) interface{}

	// Visit a parse tree produced by DmSqlParser#savepoint_name.
	VisitSavepoint_name(ctx *Savepoint_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alias_2.
	VisitAlias_2(ctx *Alias_2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#full_column_name.
	VisitFull_column_name(ctx *Full_column_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#schema_name.
	VisitSchema_name(ctx *Schema_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#not_tag.
	VisitNot_tag(ctx *Not_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#debug_tag.
	VisitDebug_tag(ctx *Debug_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#column_tag.
	VisitColumn_tag(ctx *Column_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#pendant_tag.
	VisitPendant_tag(ctx *Pendant_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#unique_tag.
	VisitUnique_tag(ctx *Unique_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#partition_tag.
	VisitPartition_tag(ctx *Partition_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#row_tag.
	VisitRow_tag(ctx *Row_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#as_tag.
	VisitAs_tag(ctx *As_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#from_tag.
	VisitFrom_tag(ctx *From_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#into_tag.
	VisitInto_tag(ctx *Into_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#work_tag.
	VisitWork_tag(ctx *Work_tagContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_grant_option.
	VisitWith_grant_option(ctx *With_grant_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_admin_option.
	VisitWith_admin_option(ctx *With_admin_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#time_zone_or_local.
	VisitTime_zone_or_local(ctx *Time_zone_or_localContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sub_plsql_datatype.
	VisitSub_plsql_datatype(ctx *Sub_plsql_datatypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#datatype_list.
	VisitDatatype_list(ctx *Datatype_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#datatype.
	VisitDatatype(ctx *DatatypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#datatype2.
	VisitDatatype2(ctx *Datatype2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#opr_dtype.
	VisitOpr_dtype(ctx *Opr_dtypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#opr_datatype_lst.
	VisitOpr_datatype_lst(ctx *Opr_datatype_lstContext) interface{}

	// Visit a parse tree produced by DmSqlParser#interval_qualifier.
	VisitInterval_qualifier(ctx *Interval_qualifierContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dtype.
	VisitDtype(ctx *DtypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#dtype1.
	VisitDtype1(ctx *Dtype1Context) interface{}

	// Visit a parse tree produced by DmSqlParser#dtype2.
	VisitDtype2(ctx *Dtype2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#double_length_option.
	VisitDouble_length_option(ctx *Double_length_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#size_unit_caluse.
	VisitSize_unit_caluse(ctx *Size_unit_caluseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#lt_integer_negative.
	VisitLt_integer_negative(ctx *Lt_integer_negativeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#create_contextindex_stmt.
	VisitCreate_contextindex_stmt(ctx *Create_contextindex_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#lexer_name.
	VisitLexer_name(ctx *Lexer_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#lexer_clause.
	VisitLexer_clause(ctx *Lexer_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#lexer_clause2.
	VisitLexer_clause2(ctx *Lexer_clause2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#sync.
	VisitSync(ctx *SyncContext) interface{}

	// Visit a parse tree produced by DmSqlParser#drop_contextindex_stmt.
	VisitDrop_contextindex_stmt(ctx *Drop_contextindex_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#alter_contextindex_stmt.
	VisitAlter_contextindex_stmt(ctx *Alter_contextindex_stmtContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cti_sync_option.
	VisitCti_sync_option(ctx *Cti_sync_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#type_name.
	VisitType_name(ctx *Type_nameContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sizeof_type.
	VisitSizeof_type(ctx *Sizeof_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#array_type.
	VisitArray_type(ctx *Array_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#builtin_types.
	VisitBuiltin_types(ctx *Builtin_typesContext) interface{}

	// Visit a parse tree produced by DmSqlParser#integral_type.
	VisitIntegral_type(ctx *Integral_typeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#sql_builtin_types.
	VisitSql_builtin_types(ctx *Sql_builtin_typesContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cursor_declaration.
	VisitCursor_declaration(ctx *Cursor_declarationContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cursor_declaration_2.
	VisitCursor_declaration_2(ctx *Cursor_declaration_2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#cursor_attrs_options.
	VisitCursor_attrs_options(ctx *Cursor_attrs_optionsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cursor_attrs.
	VisitCursor_attrs(ctx *Cursor_attrsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cursor_attr.
	VisitCursor_attr(ctx *Cursor_attrContext) interface{}

	// Visit a parse tree produced by DmSqlParser#opt_rank_specifier.
	VisitOpt_rank_specifier(ctx *Opt_rank_specifierContext) interface{}

	// Visit a parse tree produced by DmSqlParser#rank_specifiers.
	VisitRank_specifiers(ctx *Rank_specifiersContext) interface{}

	// Visit a parse tree produced by DmSqlParser#rank_specifier.
	VisitRank_specifier(ctx *Rank_specifierContext) interface{}

	// Visit a parse tree produced by DmSqlParser#opt_dim_separators.
	VisitOpt_dim_separators(ctx *Opt_dim_separatorsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#opt_rank_specifier2.
	VisitOpt_rank_specifier2(ctx *Opt_rank_specifier2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#dim_separators.
	VisitDim_separators(ctx *Dim_separatorsContext) interface{}

	// Visit a parse tree produced by DmSqlParser#opt_argument_list.
	VisitOpt_argument_list(ctx *Opt_argument_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#json_fun_tail.
	VisitJson_fun_tail(ctx *Json_fun_tailContext) interface{}

	// Visit a parse tree produced by DmSqlParser#ignore_nulls_clause.
	VisitIgnore_nulls_clause(ctx *Ignore_nulls_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mixed_param_list.
	VisitMixed_param_list(ctx *Mixed_param_listContext) interface{}

	// Visit a parse tree produced by DmSqlParser#mixed_param.
	VisitMixed_param(ctx *Mixed_paramContext) interface{}

	// Visit a parse tree produced by DmSqlParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by DmSqlParser#cursor_option.
	VisitCursor_option(ctx *Cursor_optionContext) interface{}

	// Visit a parse tree produced by DmSqlParser#without_into_select2.
	VisitWithout_into_select2(ctx *Without_into_select2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#cursor_option_2.
	VisitCursor_option_2(ctx *Cursor_option_2Context) interface{}

	// Visit a parse tree produced by DmSqlParser#region_size.
	VisitRegion_size(ctx *Region_sizeContext) interface{}

	// Visit a parse tree produced by DmSqlParser#copy_num.
	VisitCopy_num(ctx *Copy_numContext) interface{}

	// Visit a parse tree produced by DmSqlParser#redundancy_clause.
	VisitRedundancy_clause(ctx *Redundancy_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#striping_clause.
	VisitStriping_clause(ctx *Striping_clauseContext) interface{}

	// Visit a parse tree produced by DmSqlParser#with_huge_clause.
	VisitWith_huge_clause(ctx *With_huge_clauseContext) interface{}
}
