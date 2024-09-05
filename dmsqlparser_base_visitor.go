// Code generated from DmSqlParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // DmSqlParser
import "github.com/antlr4-go/antlr/v4"

type BaseDmSqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDmSqlParserVisitor) VisitDmprogram(ctx *DmprogramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSql_clauses(ctx *Sql_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDdlsql(ctx *DdlsqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDmlsql(ctx *DmlsqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPrivsql(ctx *PrivsqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOthersql(ctx *OthersqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUtilsql(ctx *UtilsqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExplainsql(ctx *ExplainsqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitShutdown_stmt(ctx *Shutdown_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_diskgroup_stmt(ctx *Alter_diskgroup_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLocal(ctx *LocalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDmsubprogram(ctx *DmsubprogramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDeclare_block(ctx *Declare_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDecl_var_cur_list_options(ctx *Decl_var_cur_list_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDecl_var_cur_list_1(ctx *Decl_var_cur_list_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDecl_var_cur_list(ctx *Decl_var_cur_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDecl_plsql_type(ctx *Decl_plsql_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPlsql_type_def(ctx *Plsql_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLt_int_lst(ctx *Lt_int_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRec_item_def_list(ctx *Rec_item_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRec_item_def(ctx *Rec_item_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDecl_variable(ctx *Decl_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNot_null(ctx *Not_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPlsql_datatype(ctx *Plsql_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDefault_clause_option(ctx *Default_clause_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVariable_name_list(ctx *Variable_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDecl_except(ctx *Decl_exceptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPragma_def(ctx *Pragma_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPragma(ctx *PragmaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPlbody(ctx *PlbodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSs_plbody(ctx *Ss_plbodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLabel_list(ctx *Label_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLabel_list_options(ctx *Label_list_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLabel_demiliter_l(ctx *Label_demiliter_lContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLabel_demiliter_r(ctx *Label_demiliter_rContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPlsql(ctx *PlsqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUr_option(ctx *Ur_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFlashback_trig_enable(ctx *Flashback_trig_enableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitScn_or_lsn(ctx *Scn_or_lsnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_table_name_list(ctx *Full_table_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFlashback_tab_stmt(ctx *Flashback_tab_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRename(ctx *RenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_system_set_stmt(ctx *Alter_system_set_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDefer(ctx *DeferContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitScope(ctx *ScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_session_stmt(ctx *Alter_session_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParallel_mode(ctx *Parallel_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParallel_degree(ctx *Parallel_degreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPurge(ctx *PurgeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSess_id(ctx *Sess_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSet_time_zone_string(ctx *Set_time_zone_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSess_attr(ctx *Sess_attrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSess_attr_val(ctx *Sess_attr_valContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSet_schema_stmt(ctx *Set_schema_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPlblock(ctx *PlblockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExcept_option(ctx *Except_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFinally_option(ctx *Finally_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFinally_tail(ctx *Finally_tailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExcept_handler_list(ctx *Except_handler_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExcept_handler(ctx *Except_handlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExcept_name(ctx *Except_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExcept_list(ctx *Except_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIf_stmt(ctx *If_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIf_stmt_clause(ctx *If_stmt_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIf_condition_clause(ctx *If_condition_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIf_then_clause(ctx *If_then_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitElseif_lst_option(ctx *Elseif_lst_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitElseif_clause(ctx *Elseif_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitElse_option(ctx *Else_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSs_if_stmt_clause(ctx *Ss_if_stmt_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSs_elseif_lst_option(ctx *Ss_elseif_lst_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSs_elseif_clause(ctx *Ss_elseif_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSs_else_option(ctx *Ss_else_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCase_stmt(ctx *Case_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPlsearched_when_clause(ctx *Plsearched_when_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPlsearched_when_list(ctx *Plsearched_when_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCase_option(ctx *Case_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAssign_stmt(ctx *Assign_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAssign_obj(ctx *Assign_objContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAssign_obj2(ctx *Assign_obj2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAssign_op(ctx *Assign_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGoto_stmt(ctx *Goto_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWhile_stmt(ctx *While_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLoop_stmt(ctx *Loop_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRepeat_stmt(ctx *Repeat_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFor_stmt(ctx *For_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitForall_stmt(ctx *Forall_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitForall_between_option(ctx *Forall_between_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitForall_save_exception_option(ctx *Forall_save_exception_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitForall_index_values(ctx *Forall_index_valuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitForall_start(ctx *Forall_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitForall_dml_stmt(ctx *Forall_dml_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIn_option(ctx *In_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFor_condition(ctx *For_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPipe_row_stmt(ctx *Pipe_row_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExit_stmt(ctx *Exit_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitContinue_stmt(ctx *Continue_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNull_stmt(ctx *Null_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPrint_stmt(ctx *Print_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExecute_stmt(ctx *Execute_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDyn_return(ctx *Dyn_returnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUsing_clause(ctx *Using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUsing_exp_list(ctx *Using_exp_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUsing_exp(ctx *Using_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_proc_stmt(ctx *Alter_proc_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_func_stmt(ctx *Alter_func_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_package_stmt(ctx *Alter_package_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPkg_type(ctx *Pkg_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDeclare_opt(ctx *Declare_optContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_table_stmt(ctx *Alter_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_tag(ctx *Alter_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_index_stmt(ctx *Alter_index_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_index_name(ctx *Full_index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_index_action(ctx *Alter_index_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRebuild_clause(ctx *Rebuild_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExclusive_options(ctx *Exclusive_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAsynchronous_options(ctx *Asynchronous_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVisible_clause(ctx *Visible_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_def_list(ctx *Column_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLock(ctx *LockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_table_partition_action(ctx *Alter_table_partition_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTemplate_info(ctx *Template_infoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTemplate_item_2(ctx *Template_item_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTemplate_item_1(ctx *Template_item_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIncluding_indexes(ctx *Including_indexesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTruncate_partition_name(ctx *Truncate_partition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCons_enable(ctx *Cons_enableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitReuse_storage_option(ctx *Reuse_storage_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_table_action(ctx *Alter_table_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFast_flag(ctx *Fast_flagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStorage_stat_flag(ctx *Storage_stat_flagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStorage_stat_cols(ctx *Storage_stat_colsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHfs_rebuild_level(ctx *Hfs_rebuild_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAta_lock_option(ctx *Ata_lock_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMdf_column_def_list(ctx *Mdf_column_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMdf_column_def(ctx *Mdf_column_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_def(ctx *Column_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_def_ex(ctx *Column_def_exContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_def_low(ctx *Column_def_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVirtual_column_datatype(ctx *Virtual_column_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVirtual_column_generated(ctx *Virtual_column_generatedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVirtual_column_virtual(ctx *Virtual_column_virtualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVirtual_column_visible(ctx *Virtual_column_visibleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVirtual_column_def(ctx *Virtual_column_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCharset_option(ctx *Charset_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_def_4_option(ctx *Column_def_4_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAuto_update_clause(ctx *Auto_update_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUpdate_exp(ctx *Update_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIdentity_clause(ctx *Identity_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDefault_clause_with_on_null_opt(ctx *Default_clause_with_on_null_optContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDefault_clause(ctx *Default_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDefault_exp(ctx *Default_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_constraint_def(ctx *Column_constraint_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitConstraint_name_def_options(ctx *Constraint_name_def_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitConstraint_name_def(ctx *Constraint_name_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_constraints(ctx *Column_constraintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_constraint(ctx *Column_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_constraint_action(ctx *Column_constraint_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNot_null_spec(ctx *Not_null_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUnique_spec(ctx *Unique_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRefs_spec(ctx *Refs_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRefs_spec_action(ctx *Refs_spec_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitForeign_key(ctx *Foreign_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRefd_table_and_columns(ctx *Refd_table_and_columnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRef_column_list(ctx *Ref_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_list(ctx *Column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_list2(ctx *Column_list2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_column_list(ctx *Full_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_list_list(ctx *Column_list_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDrop_column_list(ctx *Drop_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMatch_option(ctx *Match_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMatch_type(ctx *Match_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRef_triggered_action(ctx *Ref_triggered_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUpdate_rule(ctx *Update_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDelete_rule(ctx *Delete_ruleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRef_action(ctx *Ref_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCheck_constraint_def(ctx *Check_constraint_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCheck_condition(ctx *Check_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestrict_cascade(ctx *Restrict_cascadeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCascade_opt(ctx *Cascade_optContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitConstraint_name_options(ctx *Constraint_name_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCheck_option_def_true(ctx *Check_option_def_trueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitConstraint_attributes_options(ctx *Constraint_attributes_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitConstraint_attributes(ctx *Constraint_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDeferrable_option(ctx *Deferrable_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitConstraint_check_time(ctx *Constraint_check_timeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTable_constraint_clause(ctx *Table_constraint_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTable_constraint(ctx *Table_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUsing_index_clause(ctx *Using_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitForeign_key_clause(ctx *Foreign_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_trigger_stmt(ctx *Alter_trigger_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_trigger_option(ctx *Alter_trigger_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_table_partition_action_options(ctx *Alter_table_partition_action_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRefresh_materialized_view_stmt(ctx *Refresh_materialized_view_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitComplete_del_null(ctx *Complete_del_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRefresh_complete_del(ctx *Refresh_complete_delContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_materialized_view_stmt(ctx *Alter_materialized_view_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_view_stmt(ctx *Alter_view_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_view_action(ctx *Alter_view_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCons_novalidate(ctx *Cons_novalidateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitView_constraint_clause(ctx *View_constraint_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitView_constraint(ctx *View_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitView_unique_spec(ctx *View_unique_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitView_refs_spec(ctx *View_refs_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitView_refs_spec_action(ctx *View_refs_spec_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCall_proc_stmt(ctx *Call_proc_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRaw_call_proc_stmt(ctx *Raw_call_proc_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCall_proc_stmt_2(ctx *Call_proc_stmt_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExec_proc_stmt(ctx *Exec_proc_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDblink_clause(ctx *Dblink_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDblink_clause2(ctx *Dblink_clause2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParam_list(ctx *Param_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRaw_exp_list(ctx *Raw_exp_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExp_list_2(ctx *Exp_list_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExp_list(ctx *Exp_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIns_exp_list(ctx *Ins_exp_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLt_exp(ctx *Lt_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRange_partition_exp(ctx *Range_partition_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRange_partition_exp_list(ctx *Range_partition_exp_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitList_partition_exp(ctx *List_partition_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitList_partition_exp_list(ctx *List_partition_exp_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitList_partition_value_list(ctx *List_partition_value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitClose_cursor_stmt(ctx *Close_cursor_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitClose_cursor_statement(ctx *Close_cursor_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBegin_trans_stmt(ctx *Begin_trans_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCommit_trans_stmt(ctx *Commit_trans_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCommit_head(ctx *Commit_headContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCommit_tail(ctx *Commit_tailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCommit_wait_immed_option(ctx *Commit_wait_immed_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitConnect_stmt(ctx *Connect_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPassword(ctx *PasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTs_storage(ctx *Ts_storageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTs_storage_clause(ctx *Ts_storage_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_tablespace_stmt(ctx *Create_tablespace_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCtss_with_clause(ctx *Ctss_with_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_tablespace_set_stmt(ctx *Create_tablespace_set_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_tablespace_set_stmt(ctx *Alter_tablespace_set_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCache(ctx *CacheContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_tablespace_stmt(ctx *Alter_tablespace_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitKeep(ctx *KeepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_tablespace_action(ctx *Alter_tablespace_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFile_list(ctx *File_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPathname_list(ctx *Pathname_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInteger_list(ctx *Integer_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMirror(ctx *MirrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAutoextend_nextsize(ctx *Autoextend_nextsizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAutoextend_maxsize(ctx *Autoextend_maxsizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAutoextend(ctx *AutoextendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOn_raft(ctx *On_raftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitArchfile(ctx *ArchfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitArchflag(ctx *ArchflagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitArchstyle_options(ctx *Archstyle_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitArchstyle(ctx *ArchstyleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitArchdir(ctx *ArchdirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBakfile(ctx *BakfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParameters_option_list(ctx *Parameters_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParameter_option_list(ctx *Parameter_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParameter_option(ctx *Parameter_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPathname(ctx *PathnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPathname_options(ctx *Pathname_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_stmt(ctx *Backup_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBack_range_option(ctx *Back_range_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBack_archive_spec_null(ctx *Back_archive_spec_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNot_backed_up(ctx *Not_backed_upContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitArchive_spec(ctx *Archive_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSpec_lsn(ctx *Spec_lsnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_delete_archive(ctx *Backup_delete_archiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_options(ctx *Backup_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCumulative(ctx *CumulativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_bak_dir_list(ctx *With_bak_dir_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBase_on_backup(ctx *Base_on_backupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_to_options(ctx *Backup_to_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_path_null(ctx *Backup_path_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDevice_type(ctx *Device_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParms_command(ctx *Parms_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMedia_name(ctx *Media_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_desc_options(ctx *Backup_desc_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_desc(ctx *Backup_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_maxsize(ctx *Backup_maxsizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_limit(ctx *Backup_limitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_identified(ctx *Backup_identifiedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_compressed(ctx *Backup_compressedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_without(ctx *Backup_withoutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_tsk_thread_num_null(ctx *Backup_tsk_thread_num_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_parallel_dir(ctx *Backup_parallel_dirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_trace_file_level(ctx *Backup_trace_file_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestore_stmt(ctx *Restore_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestore_datafile_lst(ctx *Restore_datafile_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestore_mapped_file(ctx *Restore_mapped_fileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMapped_file(ctx *Mapped_fileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRes_struct(ctx *Res_structContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTsk_thread_num(ctx *Tsk_thread_numContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestore_tsk_thread_num_null(ctx *Restore_tsk_thread_num_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestore_parallel(ctx *Restore_parallelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_table_name_options(ctx *Full_table_name_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRes_without_index_constraint(ctx *Res_without_index_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestore_from(ctx *Restore_fromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRes_until(ctx *Res_untilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestore_file_list_options(ctx *Restore_file_list_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMirror_file_list_options(ctx *Mirror_file_list_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestore_trace_file_level(ctx *Restore_trace_file_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestore_file_list(ctx *Restore_file_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestore_file(ctx *Restore_fileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMirror_file_list(ctx *Mirror_file_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMirror_file(ctx *Mirror_fileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_bak_arch_dir_list(ctx *With_bak_arch_dir_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestore_identified(ctx *Restore_identifiedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_func_stmt(ctx *Create_func_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFunc_aggr_clause(ctx *Func_aggr_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPipelined_options(ctx *Pipelined_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitReplace_option(ctx *Replace_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitEdit_options(ctx *Edit_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitEncryption_option(ctx *Encryption_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCalc_option(ctx *Calc_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFunc_action(ctx *Func_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFunc_call_options(ctx *Func_call_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFunc_call_option_list(ctx *Func_call_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFunc_call_option(ctx *Func_call_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInvoker_rights_clause_options(ctx *Invoker_rights_clause_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInvoker_rights_clause(ctx *Invoker_rights_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDeterministic_clause_options(ctx *Deterministic_clause_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDeterministic_clause(ctx *Deterministic_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFunc_call_option2_options(ctx *Func_call_option2_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFunc_call_option_list2(ctx *Func_call_option_list2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFunc_call_option2(ctx *Func_call_option2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitResult_cache_clause(ctx *Result_cache_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInner_fun_name(ctx *Inner_fun_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPlatform_type(ctx *Platform_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParam_def_list_option(ctx *Param_def_list_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParam_def_list(ctx *Param_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParam_def(ctx *Param_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParam_in_out_option(ctx *Param_in_out_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIs_as(ctx *Is_asContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStat_on_org_stmt(ctx *Stat_on_org_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStat_size(ctx *Stat_sizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStat_para(ctx *Stat_paraContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStat_summarize(ctx *Stat_summarizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMstat_ex(ctx *Mstat_exContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIndexid(ctx *IndexidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGlobal_tag(ctx *Global_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBm_join_index_clause(ctx *Bm_join_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParallel_stmt(ctx *Parallel_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_index_stmt(ctx *Create_index_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_inner(ctx *With_innerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIndex_no_sort(ctx *Index_no_sortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOnline_options(ctx *Online_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUnusable_options(ctx *Unusable_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitReverse_options(ctx *Reverse_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIndex_column_list(ctx *Index_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIndex_column_name(ctx *Index_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStorage_hash_tag(ctx *Storage_hash_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStorage_hash(ctx *Storage_hashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStorage_tag(ctx *Storage_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStorage_tag_nn(ctx *Storage_tag_nnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTablespace_clause(ctx *Tablespace_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObject_table_substitution_clause(ctx *Object_table_substitution_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObject_table_substitution(ctx *Object_table_substitutionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOid_clause(ctx *Oid_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOid_gen_type(ctx *Oid_gen_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOid_index_clause(ctx *Oid_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOid_tablespace_clause(ctx *Oid_tablespace_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOid_tablespace_name(ctx *Oid_tablespace_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLocal_option(ctx *Local_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStorage_list(ctx *Storage_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStorage_hashpartmap(ctx *Storage_hashpartmapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStorage(ctx *StorageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitId_list(ctx *Id_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_proc_stmt(ctx *Create_proc_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_package_stmt(ctx *Create_package_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPkg_cls_flag(ctx *Pkg_cls_flagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBlk_end_option(ctx *Blk_end_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBlk_end_option_low(ctx *Blk_end_option_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPackage_def_list_options(ctx *Package_def_list_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPackage_def_list(ctx *Package_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPackage_def(ctx *Package_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRestrict_param_lst(ctx *Restrict_param_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_package_body_stmt(ctx *Create_package_body_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_pkg_body_header(ctx *Create_pkg_body_headerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPkg_cls_body_flag(ctx *Pkg_cls_body_flagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPackage_body_init_option(ctx *Package_body_init_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPackage_body_def_list(ctx *Package_body_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPackage_body_def(ctx *Package_body_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPackage_body_def2(ctx *Package_body_def2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCheck_exec_options(ctx *Check_exec_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpg_decl_stmt(ctx *Subpg_decl_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_type_stmt(ctx *Create_type_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitForce_option(ctx *Force_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObject_option(ctx *Object_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUnder_option(ctx *Under_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObject_def_list_options(ctx *Object_def_list_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObject_def_list(ctx *Object_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObject_def(ctx *Object_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMember_static(ctx *Member_staticContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMethod_inherit_options(ctx *Method_inherit_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMethod_inherit_option(ctx *Method_inherit_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFinal_inst_list_options(ctx *Final_inst_list_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFinal_inst_list(ctx *Final_inst_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFinal_inst(ctx *Final_instContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOverriding_option(ctx *Overriding_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMethod_attr_options(ctx *Method_attr_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMethod_attr(ctx *Method_attrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_type_body_stmt(ctx *Create_type_body_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObject_body_def_list(ctx *Object_body_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObject_body_def(ctx *Object_body_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_context_stmt(ctx *Create_context_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNamespace(ctx *NamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInitialized(ctx *InitializedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_directory_stmt(ctx *Create_directory_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_crypto_stmt(ctx *Create_crypto_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_crypto_stmt(ctx *Alter_crypto_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_crypto_action(ctx *Alter_crypto_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitComment_stmt(ctx *Comment_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitComment_tag(ctx *Comment_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_partition_group_stmt(ctx *Create_partition_group_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStorage_act_datatype(ctx *Storage_act_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPg_storage_lst(ctx *Pg_storage_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPg_storage(ctx *Pg_storageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_table_stmt(ctx *Create_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCtab_append_attr_clause(ctx *Ctab_append_attr_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCtab_append_attr_list(ctx *Ctab_append_attr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCobjtab_append_attr_clause(ctx *Cobjtab_append_attr_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCobjtab_append_attr_list(ctx *Cobjtab_append_attr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCtab_append_attr(ctx *Ctab_append_attrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCobjtab_append_attr(ctx *Cobjtab_append_attrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_table_action(ctx *Create_table_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCtab_log_options(ctx *Ctab_log_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCtab_log_option(ctx *Ctab_log_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCtab_error_options(ctx *Ctab_error_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAdvance_log_clause(ctx *Advance_log_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAdd_log_clause(ctx *Add_log_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCtab_error_option(ctx *Ctab_error_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCtab_row_movement_clause(ctx *Ctab_row_movement_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRange_distribute_act(ctx *Range_distribute_actContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRange_distribute_act_lst(ctx *Range_distribute_act_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitList_distribute_act(ctx *List_distribute_actContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitList_distribute_act_list(ctx *List_distribute_act_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDistribute_by_option(ctx *Distribute_by_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDistribute_by(ctx *Distribute_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIncrement_set(ctx *Increment_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIncrement(ctx *IncrementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRowdependencies_clause(ctx *Rowdependencies_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPg_sub_partition(ctx *Pg_sub_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTable_type_option(ctx *Table_type_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTable_temp_option(ctx *Table_temp_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObjtab_elem_constraint(ctx *Objtab_elem_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObjtab_element_cons_list(ctx *Objtab_element_cons_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObjtab_element_cons(ctx *Objtab_element_consContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitObjcol_constraint(ctx *Objcol_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTable_element_list(ctx *Table_element_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTable_element(ctx *Table_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTable_constraint_def(ctx *Table_constraint_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOn_commit_option(ctx *On_commit_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOn_commit_option_nn(ctx *On_commit_option_nnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLogging_option(ctx *Logging_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLogging_option_nn(ctx *Logging_option_nnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPartition_clause(ctx *Partition_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPartition_clause_nn(ctx *Partition_clause_nnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHorizon_partition_clause(ctx *Horizon_partition_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCompress_tag_hdr(ctx *Compress_tag_hdrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCompress_clause_opt(ctx *Compress_clause_optContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCompress_tag(ctx *Compress_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCompress_level(ctx *Compress_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCompress_type(ctx *Compress_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRange_partition(ctx *Range_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRange_partition_list(ctx *Range_partition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHash_partition(ctx *Hash_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHash_partition_list(ctx *Hash_partition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitList_partition(ctx *List_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitList_partition_list(ctx *List_partition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSplit_partition_list(ctx *Split_partition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPartition_act(ctx *Partition_actContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVertical_partition_act(ctx *Vertical_partition_actContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInterval_item(ctx *Interval_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHorizon_partition_act_datatype(ctx *Horizon_partition_act_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHorizon_partition_act(ctx *Horizon_partition_actContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLock_partitions_clause(ctx *Lock_partitions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartion_template_list_datatype_options(ctx *Subpartion_template_list_datatype_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartion_template_list_datatype(ctx *Subpartion_template_list_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartion_template_list_options(ctx *Subpartion_template_list_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartion_template_list(ctx *Subpartion_template_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartion_template_datatype(ctx *Subpartion_template_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRange_subpartion_template_datatype(ctx *Range_subpartion_template_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHash_subpartion_template_datatype(ctx *Hash_subpartion_template_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHash_subpartions_template_datatype_option(ctx *Hash_subpartions_template_datatype_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitList_subpartion_template_datatype(ctx *List_subpartion_template_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartion_template(ctx *Subpartion_templateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRange_subpartion_template(ctx *Range_subpartion_templateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHash_subpartion_template(ctx *Hash_subpartion_templateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHash_subpartions_template_option(ctx *Hash_subpartions_template_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitList_subpartion_template(ctx *List_subpartion_templateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRange_subpartition(ctx *Range_subpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHash_subpartition(ctx *Hash_subpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitList_subpartition(ctx *List_subpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRange_subpartition_list(ctx *Range_subpartition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHash_subpartition_list(ctx *Hash_subpartition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitList_subpartition_list(ctx *List_subpartition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartition_hash_desc(ctx *Subpartition_hash_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartition_range_desc(ctx *Subpartition_range_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartition_list_desc(ctx *Subpartition_list_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartition_hash_desc_list(ctx *Subpartition_hash_desc_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartition_range_desc_list(ctx *Subpartition_range_desc_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartition_list_desc_list(ctx *Subpartition_list_desc_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartition_desc(ctx *Subpartition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSubpartition_desc_option(ctx *Subpartition_desc_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAdd_subpartition_desc(ctx *Add_subpartition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPartition_no(ctx *Partition_noContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitComment_clause(ctx *Comment_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitEncrypt_clause_options(ctx *Encrypt_clause_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitEncrypt_clause(ctx *Encrypt_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitEncrypt_cipher(ctx *Encrypt_cipherContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCrypto_name(ctx *Crypto_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCipher_name(ctx *Cipher_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_cipher_name(ctx *Full_cipher_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitEncrypt_type(ctx *Encrypt_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitManual_clause(ctx *Manual_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUser_clause_options(ctx *User_clause_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUser_clause(ctx *User_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUser_list_option(ctx *User_list_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUser_list(ctx *User_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHash_cipher(ctx *Hash_cipherContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHash_type(ctx *Hash_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSpace_limit(ctx *Space_limitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSpace_limit_nn(ctx *Space_limit_nnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSpace_limit_1(ctx *Space_limit_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSpace_limit2(ctx *Space_limit2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDel_res(ctx *Del_resContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_enable(ctx *Trig_enableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAt_raft(ctx *At_raftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_trigger_stmt(ctx *Create_trigger_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBefore_after(ctx *Before_afterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_del_ins_upd_list(ctx *Trig_del_ins_upd_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_del_ins_upd(ctx *Trig_del_ins_updContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUpdate_of_option(ctx *Update_of_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNowait(ctx *NowaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_event_list(ctx *Trig_event_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_event(ctx *Trig_eventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitEvent_object(ctx *Event_objectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_referencing_def_options(ctx *Trig_referencing_def_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_referencing_def(ctx *Trig_referencing_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_referencing_list(ctx *Trig_referencing_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_referencing_old(ctx *Trig_referencing_oldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_referencing_new(ctx *Trig_referencing_newContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_for_each_option(ctx *Trig_for_each_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_timer_rate(ctx *Trig_timer_rateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExec_ep_seqno(ctx *Exec_ep_seqnoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRate_over_day(ctx *Rate_over_dayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMonth_rate(ctx *Month_rateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDay_in_month(ctx *Day_in_monthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDay_in_month_week(ctx *Day_in_month_weekContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWeek_rate(ctx *Week_rateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDay_of_week_list(ctx *Day_of_week_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDay_rate(ctx *Day_rateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRate_in_day(ctx *Rate_in_dayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOnce_in_day(ctx *Once_in_dayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTimes_in_day(ctx *Times_in_dayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDuaring_time(ctx *Duaring_timeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDuaring_date(ctx *Duaring_dateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_when_option(ctx *Trig_when_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_when_condition(ctx *Trig_when_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRepeat_interval_stmt(ctx *Repeat_interval_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMax_run_duration(ctx *Max_run_durationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRepeat_interval(ctx *Repeat_intervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFrequency_clause(ctx *Frequency_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFrequency_define(ctx *Frequency_defineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPredefined_frequency(ctx *Predefined_frequencyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInterval_clause_list(ctx *Interval_clause_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInterval_clause_single(ctx *Interval_clause_singleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInterval_clause(ctx *Interval_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIntervalnum(ctx *IntervalnumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBymonth_clause(ctx *Bymonth_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMonthlist(ctx *MonthlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMonth(ctx *MonthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNumeric_month(ctx *Numeric_monthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitChar_month(ctx *Char_monthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitByweekno_clause(ctx *Byweekno_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWeekno_list(ctx *Weekno_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWeekno(ctx *WeeknoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitByyearday_clause(ctx *Byyearday_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitYearday_list(ctx *Yearday_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitYearday(ctx *YeardayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBymonthday_clause(ctx *Bymonthday_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMonthday_list(ctx *Monthday_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMonthday(ctx *MonthdayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitByday_clause(ctx *Byday_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitByday_list(ctx *Byday_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitByday(ctx *BydayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWeekdaynum_options(ctx *Weekdaynum_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWeekdaynum(ctx *WeekdaynumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDay(ctx *DayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitByhour_clause(ctx *Byhour_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHour_list(ctx *Hour_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHour(ctx *HourContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitByminute_clause(ctx *Byminute_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMinute_list(ctx *Minute_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMinute(ctx *MinuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBysecond_clause(ctx *Bysecond_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSecond_list(ctx *Second_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSecond(ctx *SecondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQuery_rewrite(ctx *Query_rewriteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBuild_clause(ctx *Build_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMv_refresh_option(ctx *Mv_refresh_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMv_refresh_option_list(ctx *Mv_refresh_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMv_refresh_clause(ctx *Mv_refresh_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMv_log_purge_syn_asyn_clause(ctx *Mv_log_purge_syn_asyn_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMv_log_purge_clause(ctx *Mv_log_purge_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMv_log_with_syntax_item(ctx *Mv_log_with_syntax_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMv_log_with_syntax_item_list(ctx *Mv_log_with_syntax_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMv_log_including_new_values(ctx *Mv_log_including_new_valuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMv_log_with_clause_null(ctx *Mv_log_with_clause_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_materialized_view_log_stmt(ctx *Create_materialized_view_log_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPrebuilt_table_clause_null(ctx *Prebuilt_table_clause_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_materialized_view_stmt(ctx *Create_materialized_view_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_view_stmt(ctx *Create_view_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_view_stmt_body(ctx *Create_view_stmt_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_list3_options(ctx *Column_list3_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_list3(ctx *Column_list3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitView_column_constraint_def(ctx *View_column_constraint_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitView_column_constraints(ctx *View_column_constraintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitView_column_constraint(ctx *View_column_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitView_column_constraint_action(ctx *View_column_constraint_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitView_constraint_def(ctx *View_constraint_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_schemabinding(ctx *With_schemabindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_list_options(ctx *Column_list_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_check_or_readonly_option(ctx *With_check_or_readonly_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCheck_level_option(ctx *Check_level_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDecl_cursor(ctx *Decl_cursorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDelete_stmt(ctx *Delete_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDelete_stmt_body(ctx *Delete_stmt_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDelete_multi_tv_option(ctx *Delete_multi_tv_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCheck_limit_option(ctx *Check_limit_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWhere_current_option(ctx *Where_current_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWhere_clause(ctx *Where_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStart_with_clause_null(ctx *Start_with_clause_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitConnect_by_item(ctx *Connect_by_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitConnect_by_clause(ctx *Connect_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHierarchical_query_clause(ctx *Hierarchical_query_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNocycle_flag(ctx *Nocycle_flagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSearch_condition(ctx *Search_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDisconnect_stmt(ctx *Disconnect_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDisconnect_option(ctx *Disconnect_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDrop_stmt(ctx *Drop_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDrop_stmt_body_1(ctx *Drop_stmt_body_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDrop_stmt_2(ctx *Drop_stmt_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDrop_id_db_object(ctx *Drop_id_db_objectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitId_db_object(ctx *Id_db_objectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDrop_db_object(ctx *Drop_db_objectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExist(ctx *ExistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNot_exist(ctx *Not_existContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDb_object(ctx *Db_objectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIs_detach(ctx *Is_detachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPurge_option(ctx *Purge_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_database_stmt(ctx *Alter_database_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_system_action(ctx *Alter_system_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_database_action(ctx *Alter_database_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitForce(ctx *ForceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTablespace_name(ctx *Tablespace_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRaft_name(ctx *Raft_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFetch_into(ctx *Fetch_intoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBulk_or_single_into(ctx *Bulk_or_single_intoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFetch_stmt(ctx *Fetch_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFetch_statement(ctx *Fetch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFetch_tail(ctx *Fetch_tailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFetch_limit_option(ctx *Fetch_limit_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFetch_option(ctx *Fetch_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFetch_direct_option(ctx *Fetch_direct_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLog_errors_into(ctx *Log_errors_intoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLog_errors_expression(ctx *Log_errors_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLog_errors_unlimited(ctx *Log_errors_unlimitedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLog_errors(ctx *Log_errorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInsert_stmt(ctx *Insert_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInsert_stmt_body(ctx *Insert_stmt_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_column_list_options(ctx *Full_column_list_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIns_value_options(ctx *Ins_value_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInsert_into_single(ctx *Insert_into_singleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMulti_insert_into_list(ctx *Multi_insert_into_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMulti_insert_tag(ctx *Multi_insert_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInsert_into_single_condition(ctx *Insert_into_single_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMulti_insert_into_condition_list(ctx *Multi_insert_into_condition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMulti_insert_into_else(ctx *Multi_insert_into_elseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMulti_insert_stmt_body(ctx *Multi_insert_stmt_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInsert_tail(ctx *Insert_tailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInsert_action(ctx *Insert_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRecord_var_values(ctx *Record_var_valuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRecord_var_value(ctx *Record_var_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIns_value(ctx *Ins_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOpen_stmt(ctx *Open_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOpen_statement(ctx *Open_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOpen_tail(ctx *Open_tailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitReturn_stmt(ctx *Return_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRaise_stmt(ctx *Raise_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRollback_stmt(ctx *Rollback_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTo_savepoint(ctx *To_savepointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSavepoint_stmt(ctx *Savepoint_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSelect_stmt(ctx *Select_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAll_distinct_option(ctx *All_distinct_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAll_distinct_option_2(ctx *All_distinct_option_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCorresponding_clause(ctx *Corresponding_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTop_option(ctx *Top_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLimit_option(ctx *Limit_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLimit_clause(ctx *Limit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLimit_not_empty(ctx *Limit_not_emptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRow_limiting_clause(ctx *Row_limiting_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRow_num_desc(ctx *Row_num_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFirst_next_desc(ctx *First_next_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSelect_item_list(ctx *Select_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSelect_item(ctx *Select_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAs_alias(ctx *As_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSelect_tail(ctx *Select_tailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFrom_clause(ctx *From_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFrom_tv_list(ctx *From_tv_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFrom_tv(ctx *From_tvContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJoined_table(ctx *Joined_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrxid(ctx *TrxidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFlashback_query_low(ctx *Flashback_query_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrxid_option(ctx *Trxid_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRange_from_to(ctx *Range_from_toContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSample_exp(ctx *Sample_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPivot_sfun(ctx *Pivot_sfunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPivot_sfun_lst(ctx *Pivot_sfun_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPivot_for_clause(ctx *Pivot_for_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPivot_in_clause1_expr(ctx *Pivot_in_clause1_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPivot_in_clause_low_1(ctx *Pivot_in_clause_low_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPivot_in_clause_low_2(ctx *Pivot_in_clause_low_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPivot_in_clause_low(ctx *Pivot_in_clause_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPivot_xml(ctx *Pivot_xmlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPivot_clause_low(ctx *Pivot_clause_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUnpivot_val_col_lst(ctx *Unpivot_val_col_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInclude_clause(ctx *Include_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUnpivot_in_clause_expr(ctx *Unpivot_in_clause_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUnpivot_in_clause_low(ctx *Unpivot_in_clause_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUnpivot_clause_low(ctx *Unpivot_clause_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSample_clause_low(ctx *Sample_clause_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNormal_tv_name(ctx *Normal_tv_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNormal_tv_tail(ctx *Normal_tv_tailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNormal_tv_tail_low(ctx *Normal_tv_tail_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNormal_alias(ctx *Normal_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNormal_tv_tail_low2(ctx *Normal_tv_tail_low2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNormal_tv_tail_low3(ctx *Normal_tv_tail_low3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNormal_tv_derived_table_options(ctx *Normal_tv_derived_table_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNormal_tv_derived_table_low(ctx *Normal_tv_derived_table_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNormal_tv_derived_table(ctx *Normal_tv_derived_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSelect_with_paran_with_alias(ctx *Select_with_paran_with_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFrom_table_exp(ctx *From_table_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFrom_table_select_with_paran(ctx *From_table_select_with_paranContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNormal_tv(ctx *Normal_tvContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitXml_passing(ctx *Xml_passingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitXmlcoldef_lst_options(ctx *Xmlcoldef_lst_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitXmlcoldef_lst(ctx *Xmlcoldef_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitXmlcoldef(ctx *XmlcoldefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOn_error(ctx *On_errorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJsoncol_lst(ctx *Jsoncol_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJsoncoldef_lst(ctx *Jsoncoldef_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJsoncoldef(ctx *JsoncoldefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJson_exists_col(ctx *Json_exists_colContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJson_qurey_col(ctx *Json_qurey_colContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJson_value_col(ctx *Json_value_colContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJson_nested_col(ctx *Json_nested_colContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOrdinality_col(ctx *Ordinality_colContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJoined_table_element(ctx *Joined_table_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCross_outer_apply_clause(ctx *Cross_outer_apply_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCross_outer_apply_join(ctx *Cross_outer_apply_joinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCross_join(ctx *Cross_joinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPartition_out_join(ctx *Partition_out_joinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQualified_join(ctx *Qualified_joinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQualified_joinspec(ctx *Qualified_joinspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNamed_columns_join(ctx *Named_columns_joinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJoin_hint(ctx *Join_hintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJoin_type(ctx *Join_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOuter_join_type(ctx *Outer_join_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJoin_condition(ctx *Join_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGroup_by_clause(ctx *Group_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGroup_item(ctx *Group_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExp_rollup_cube_item2(ctx *Exp_rollup_cube_item2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExp_rollup_cube_item(ctx *Exp_rollup_cube_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGrouping_set_items(ctx *Grouping_set_itemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGrouping_set_item(ctx *Grouping_set_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHaving_clause(ctx *Having_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWithout_into_select(ctx *Without_into_selectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSel_clause_app(ctx *Sel_clause_appContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSelect_clause(ctx *Select_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSimple_select(ctx *Simple_selectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSelect_with_paran(ctx *Select_with_paranContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQuery_exp(ctx *Query_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFor_xml_path(ctx *For_xml_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_tag(ctx *With_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_option(ctx *With_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_clause(ctx *With_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_function_list(ctx *With_function_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFunc_def_inner(ctx *Func_def_innerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_function_list_element(ctx *With_function_list_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_view_list(ctx *With_view_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDepth_type_option(ctx *Depth_type_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSearch_clause(ctx *Search_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCycle_clause(ctx *Cycle_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_view_list_element(ctx *With_view_list_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAssignment_obj_list(ctx *Assignment_obj_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAssignment_obj(ctx *Assignment_objContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOrder_by_options(ctx *Order_by_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOrder_by(ctx *Order_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAsc_desc_option(ctx *Asc_desc_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNulls_last_option(ctx *Nulls_last_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCollate_option(ctx *Collate_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOrder_by_list(ctx *Order_by_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOrder_by_item(ctx *Order_by_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFor_update_options(ctx *For_update_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFor_update(ctx *For_updateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSet_session_stmt(ctx *Set_session_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSet_trans_stmt(ctx *Set_trans_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrans_mode_lstl(ctx *Trans_mode_lstlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrans_mode_lst(ctx *Trans_mode_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrans_mode(ctx *Trans_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTime_zone_exp_new(ctx *Time_zone_exp_newContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrans_rw_option(ctx *Trans_rw_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrans_level_option(ctx *Trans_level_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLock_table_stmt(ctx *Lock_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLock_mode_option(ctx *Lock_mode_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSet_identins_stmt(ctx *Set_identins_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSet_identins_option(ctx *Set_identins_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrunc_table_stmt(ctx *Trunc_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUpdate_stmt(ctx *Update_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUpdate_stmt_body(ctx *Update_stmt_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUpdate_tv_list(ctx *Update_tv_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitReturn_item(ctx *Return_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitReturn_item_list(ctx *Return_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitReturn_option(ctx *Return_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitReturn_into_obj(ctx *Return_into_objContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCollect_into_rset(ctx *Collect_into_rsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlias_option(ctx *Alias_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSet_value_list(ctx *Set_value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSet_value_node(ctx *Set_value_nodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSet_col_list(ctx *Set_col_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUpdate_from_clause(ctx *Update_from_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMerge_into_stmt(ctx *Merge_into_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMerge_into_stmt_body(ctx *Merge_into_stmt_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMerge_into_sub_clause(ctx *Merge_into_sub_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMerge_update_clause(ctx *Merge_update_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMerge_insert_clause(ctx *Merge_insert_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_profile_stmt(ctx *Create_profile_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_profile_stmt(ctx *Alter_profile_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_user_stmt(ctx *Create_user_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDefault_ts_name(ctx *Default_ts_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDefault_ts_name_lst(ctx *Default_ts_name_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDefault_ts_name_node(ctx *Default_ts_name_nodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDefault_idx_ts_name(ctx *Default_idx_ts_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDefault_ts_name_low(ctx *Default_ts_name_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTemp_ts_name(ctx *Temp_ts_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDefault_ts_group_name_low(ctx *Default_ts_group_name_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOn_schema(ctx *On_schemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitReplace_old_pwd(ctx *Replace_old_pwdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_user_stmt(ctx *Alter_user_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUser_encrypt_options(ctx *User_encrypt_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUser_encrypt_option(ctx *User_encrypt_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAuthent_type_options(ctx *Authent_type_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitHash_cipher_option(ctx *Hash_cipher_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAuthent_type(ctx *Authent_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitForce_format(ctx *Force_formatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAs(ctx *AsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPwd_policy(ctx *Pwd_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAccount_lock(ctx *Account_lockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRead_only_flag(ctx *Read_only_flagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRead_only_flag_not_null(ctx *Read_only_flag_not_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitResource(ctx *ResourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExpire(ctx *ExpireContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitResource_limit_options(ctx *Resource_limit_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitResource_limit_list(ctx *Resource_limit_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitResource_limit_list_with_comma(ctx *Resource_limit_list_with_commaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitResource_limit_list_with_empty(ctx *Resource_limit_list_with_emptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitResource_limit(ctx *Resource_limitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitResource_limit_value(ctx *Resource_limit_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_audit_rule_stmt(ctx *Create_audit_rule_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRule_name(ctx *Rule_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAudit_rule_action(ctx *Audit_rule_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBy_login_or_all(ctx *By_login_or_allContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAllow_ip_list(ctx *Allow_ip_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNot_allow_ip_list(ctx *Not_allow_ip_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIp_list(ctx *Ip_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAllow_dt_list(ctx *Allow_dt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNot_allow_dt_list(ctx *Not_allow_dt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDt_list(ctx *Dt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDt(ctx *DtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOp_freq(ctx *Op_freqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQuota_val(ctx *Quota_valContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQuota_ts_node(ctx *Quota_ts_nodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQuota_ts_lst(ctx *Quota_ts_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQuota_ts(ctx *Quota_tsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_role_stmt(ctx *Create_role_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_dblink_stmt(ctx *Create_dblink_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDb_type_str(ctx *Db_type_strContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDblink_option_lst_options(ctx *Dblink_option_lst_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDblink_option_lst(ctx *Dblink_option_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDblink_option(ctx *Dblink_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_synonym_stmt(ctx *Create_synonym_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_synonym_name(ctx *Full_synonym_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_obj_name(ctx *Full_obj_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_domain_stmt(ctx *Create_domain_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDomain_default_option(ctx *Domain_default_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDomain_constraints_option(ctx *Domain_constraints_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDomain_constraints_def(ctx *Domain_constraints_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDomain_constraints(ctx *Domain_constraintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDomain_constraint(ctx *Domain_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDomain_constraint_name_def_options(ctx *Domain_constraint_name_def_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDomain_constraint_name_def(ctx *Domain_constraint_name_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDomain_constraint_name(ctx *Domain_constraint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_character_set_stmt(ctx *Create_character_set_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCharacter_set_source(ctx *Character_set_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExisting_character_set_name(ctx *Existing_character_set_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCharacter_set_name(ctx *Character_set_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCollate_clause_option(ctx *Collate_clause_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCollation_name(ctx *Collation_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_collation_stmt(ctx *Create_collation_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExisting_collation_name(ctx *Existing_collation_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPad_option(ctx *Pad_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_sequence_stmt(ctx *Create_sequence_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSequence_option_list_options(ctx *Sequence_option_list_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSequence_option_list(ctx *Sequence_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSequence_option(ctx *Sequence_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSequence_name(ctx *Sequence_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIncrement_option(ctx *Increment_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStart_option(ctx *Start_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCurrent_option(ctx *Current_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMaxvalue_option(ctx *Maxvalue_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMinvalue_option(ctx *Minvalue_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCycle_option(ctx *Cycle_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCache_option(ctx *Cache_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOrder_option(ctx *Order_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSeq_local_option(ctx *Seq_local_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWhenever_stmt_options(ctx *Whenever_stmt_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWhenever_stmt(ctx *Whenever_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGrant_stmt(ctx *Grant_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGrant_tag(ctx *Grant_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGrant_stmt_body(ctx *Grant_stmt_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRevoke_stmt(ctx *Revoke_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRevoke_stmt_body(ctx *Revoke_stmt_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGrant_privs(ctx *Grant_privsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGrant_priv_list(ctx *Grant_priv_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGrant_priv_off(ctx *Grant_priv_offContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGrant_priv(ctx *Grant_privContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRevoke_action(ctx *Revoke_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDb_priv_list(ctx *Db_priv_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDb_priv(ctx *Db_privContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBy_grantor(ctx *By_grantorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGrantees(ctx *GranteesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_schema_stmt(ctx *Create_schema_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOprt_arg(ctx *Oprt_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOprt_arg_lst(ctx *Oprt_arg_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_operator_stmt(ctx *Create_operator_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDrop_operator_stmt(ctx *Drop_operator_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGrant_and_ddl(ctx *Grant_and_ddlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTop_exp(ctx *Top_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitU_oprt(ctx *U_oprtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQualified_u_oprt(ctx *Qualified_u_oprtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExp_u_oprt(ctx *Exp_u_oprtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRaw_exp(ctx *Raw_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitExp(ctx *ExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFrom_first_last_option(ctx *From_first_last_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAfun_arg_lst(ctx *Afun_arg_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAfun_arg_lst_low(ctx *Afun_arg_lst_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIn_value_exp(ctx *In_value_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAfun_partition_by(ctx *Afun_partition_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAfun_windowing(ctx *Afun_windowingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAfun_windowing_type(ctx *Afun_windowing_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAfun_range_clause(ctx *Afun_range_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPexp(ctx *PexpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPexp_pfx(ctx *Pexp_pfxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPexp_cast(ctx *Pexp_castContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPexp_b(ctx *Pexp_bContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPexp_a(ctx *Pexp_aContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPexp_c(ctx *Pexp_cContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPexp_c_insert(ctx *Pexp_c_insertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPexp_d(ctx *Pexp_dContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPexp_e(ctx *Pexp_eContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPexp_pfx2(ctx *Pexp_pfx2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMember2(ctx *Member2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPexp_c2_insert(ctx *Pexp_c2_insertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMember_access2(ctx *Member_access2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInvocation_expression2(ctx *Invocation_expression2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMember(ctx *MemberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMember_access(ctx *Member_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInvocation_expression(ctx *Invocation_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInvocation_expression_low(ctx *Invocation_expression_lowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitXmlagg_inv_expression(ctx *Xmlagg_inv_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitXmlfun_inv_expression(ctx *Xmlfun_inv_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitXmlele_name(ctx *Xmlele_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitXmlele_sub_lst(ctx *Xmlele_sub_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitXmlattr_lst(ctx *Xmlattr_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitXmlattr(ctx *XmlattrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitXmlval_lst(ctx *Xmlval_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitKeep_clause(ctx *Keep_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWithin_clause(ctx *Within_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTypeof_expression(ctx *Typeof_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNew_obj_expression(ctx *New_obj_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNew_arr_expression(ctx *New_arr_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitArray_creation_expression(ctx *Array_creation_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPlsql_datatype_ex(ctx *Plsql_datatype_exContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNew_array_type(ctx *New_array_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOpt_array_initializer(ctx *Opt_array_initializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitArray_initializer(ctx *Array_initializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVariable_initializer_list(ctx *Variable_initializer_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVariable_initializer(ctx *Variable_initializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOpt_comma(ctx *Opt_commaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSizeof_expression(ctx *Sizeof_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitElement_access(ctx *Element_accessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDecode_case(ctx *Decode_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitElse_exp(ctx *Else_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBoolean_case(ctx *Boolean_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIf_exp(ctx *If_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBool_when_list(ctx *Bool_when_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOps(ctx *OpsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitValue_list(ctx *Value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIn_value_list(ctx *In_value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitValue_list_set(ctx *Value_list_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitComma_list(ctx *Comma_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIns_value_list(ctx *Ins_value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNull_value(ctx *Null_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitId_and_rsvd_word_others(ctx *Id_and_rsvd_word_othersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitId_and_rsvd_word(ctx *Id_and_rsvd_wordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStm_param(ctx *Stm_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStm_param_normal(ctx *Stm_param_normalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStm_param_name(ctx *Stm_param_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParam_name_options(ctx *Param_name_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitContains_query_exp(ctx *Contains_query_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitContains_query_exp_lst(ctx *Contains_query_exp_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitContains_exp(ctx *Contains_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStrict_option(ctx *Strict_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_unique_option(ctx *With_unique_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitType_option(ctx *Type_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitType_element(ctx *Type_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitType_element_list(ctx *Type_element_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBool_exp(ctx *Bool_expContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBool_exp_element(ctx *Bool_exp_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQuery_any_options(ctx *Query_any_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitGlobal_var(ctx *Global_varContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitReserved_word(ctx *Reserved_wordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNew_none_reserved_word(ctx *New_none_reserved_wordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInterval_nresvd_word(ctx *Interval_nresvd_wordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVariable_resvd_word(ctx *Variable_resvd_wordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlias_resvd_word(ctx *Alias_resvd_wordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSchname_resvd_word(ctx *Schname_resvd_wordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRaw_id(ctx *Raw_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQualified_name(ctx *Qualified_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitQualified_name2(ctx *Qualified_name2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitVariable_name(ctx *Variable_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitEnd_loop_label_null(ctx *End_loop_label_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLabel_name_options(ctx *Label_name_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLabel_name(ctx *Label_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDatabase_name(ctx *Database_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBackup_name(ctx *Backup_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_proc_name(ctx *Full_proc_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_proc_name2(ctx *Full_proc_name2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_fun_name(ctx *Full_fun_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_table_name(ctx *Full_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_grp_name(ctx *Full_grp_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_table_name2(ctx *Full_table_name2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_partition_name(ctx *Full_partition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_schema_name(ctx *Full_schema_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_name(ctx *Column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitConstraint_name(ctx *Constraint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_trigger_name(ctx *Full_trigger_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_trigger_name2(ctx *Full_trigger_name2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_view_name(ctx *Full_view_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_view_name2(ctx *Full_view_name2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCursor_name(ctx *Cursor_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrigger_name(ctx *Trigger_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLogin_name(ctx *Login_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitProfile_name(ctx *Profile_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUser_name(ctx *User_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRole_name(ctx *Role_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUser_role_name(ctx *User_role_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRole_name_list(ctx *Role_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_func_name(ctx *Full_func_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitParam_name(ctx *Param_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIndex_name(ctx *Index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIndex_name2(ctx *Index_name2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_old_name(ctx *Trig_old_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTrig_new_name(ctx *Trig_new_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_tv_name(ctx *Full_tv_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_object_name(ctx *Full_object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOrient_option(ctx *Orient_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDatepart(ctx *DatepartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDatepart_op(ctx *Datepart_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDatead_fun(ctx *Datead_funContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitReturning(ctx *ReturningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPretty(ctx *PrettyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWrapper_flag(ctx *Wrapper_flagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitArray_wrapper(ctx *Array_wrapperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJson_tail_on_empty(ctx *Json_tail_on_emptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitEmpty_handle(ctx *Empty_handleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJson_tail_on_error_null(ctx *Json_tail_on_error_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitError_handle(ctx *Error_handleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSavepoint_name(ctx *Savepoint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlias_2(ctx *Alias_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFull_column_name(ctx *Full_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSchema_name(ctx *Schema_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitNot_tag(ctx *Not_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDebug_tag(ctx *Debug_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitColumn_tag(ctx *Column_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPendant_tag(ctx *Pendant_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitUnique_tag(ctx *Unique_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitPartition_tag(ctx *Partition_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRow_tag(ctx *Row_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAs_tag(ctx *As_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitFrom_tag(ctx *From_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInto_tag(ctx *Into_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWork_tag(ctx *Work_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_grant_option(ctx *With_grant_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_admin_option(ctx *With_admin_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitTime_zone_or_local(ctx *Time_zone_or_localContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSub_plsql_datatype(ctx *Sub_plsql_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDatatype_list(ctx *Datatype_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDatatype(ctx *DatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDatatype2(ctx *Datatype2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOpr_dtype(ctx *Opr_dtypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOpr_datatype_lst(ctx *Opr_datatype_lstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitInterval_qualifier(ctx *Interval_qualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDtype(ctx *DtypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDtype1(ctx *Dtype1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDtype2(ctx *Dtype2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDouble_length_option(ctx *Double_length_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSize_unit_caluse(ctx *Size_unit_caluseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLt_integer_negative(ctx *Lt_integer_negativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCreate_contextindex_stmt(ctx *Create_contextindex_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLexer_name(ctx *Lexer_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLexer_clause(ctx *Lexer_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitLexer_clause2(ctx *Lexer_clause2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSync(ctx *SyncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDrop_contextindex_stmt(ctx *Drop_contextindex_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitAlter_contextindex_stmt(ctx *Alter_contextindex_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCti_sync_option(ctx *Cti_sync_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitType_name(ctx *Type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSizeof_type(ctx *Sizeof_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitArray_type(ctx *Array_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitBuiltin_types(ctx *Builtin_typesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIntegral_type(ctx *Integral_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitSql_builtin_types(ctx *Sql_builtin_typesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCursor_declaration(ctx *Cursor_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCursor_declaration_2(ctx *Cursor_declaration_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCursor_attrs_options(ctx *Cursor_attrs_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCursor_attrs(ctx *Cursor_attrsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCursor_attr(ctx *Cursor_attrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOpt_rank_specifier(ctx *Opt_rank_specifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRank_specifiers(ctx *Rank_specifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRank_specifier(ctx *Rank_specifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOpt_dim_separators(ctx *Opt_dim_separatorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOpt_rank_specifier2(ctx *Opt_rank_specifier2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitDim_separators(ctx *Dim_separatorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitOpt_argument_list(ctx *Opt_argument_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitJson_fun_tail(ctx *Json_fun_tailContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitIgnore_nulls_clause(ctx *Ignore_nulls_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMixed_param_list(ctx *Mixed_param_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitMixed_param(ctx *Mixed_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCursor_option(ctx *Cursor_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWithout_into_select2(ctx *Without_into_select2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCursor_option_2(ctx *Cursor_option_2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRegion_size(ctx *Region_sizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitCopy_num(ctx *Copy_numContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitRedundancy_clause(ctx *Redundancy_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitStriping_clause(ctx *Striping_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDmSqlParserVisitor) VisitWith_huge_clause(ctx *With_huge_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}
