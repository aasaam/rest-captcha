#!/usr/bin/python3
import os
import subprocess
import datetime
from pathlib import Path

def execute(command):
  process = subprocess.Popen(command, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True)
  output, _ = process.communicate()
  return output.decode("utf-8")

project_path = os.path.dirname(os.path.realpath(__file__))

fonts_source_dir = project_path + "/captcha-fonts/source"
fonts_destination_dir = project_path + "/captcha-fonts/destination"

unicode_chars = 'U+0030,U+0031,U+0032,U+0033,U+0034,U+0035,U+0036,U+0037,U+0038,U+0039,U+0660,U+0661,U+0662,U+0663,U+0664,U+0665,U+0666,U+0667,U+0668,U+0669,U+06F0,U+06F1,U+06F2,U+06F3,U+06F4,U+06F5,U+06F6,U+06F7,U+06F8,U+06F9,U+0966,U+0967,U+0968,U+0969,U+096A,U+096B,U+096C,U+096D,U+096E,U+096F,U+09E6,U+09E7,U+09E8,U+09E9,U+09EA,U+09EB,U+09EC,U+09ED,U+09EE,U+09EF,U+0F20,U+0F21,U+0F22,U+0F23,U+0F24,U+0F25,U+0F26,U+0F27,U+0F28,U+0F29,U+1040,U+1041,U+1042,U+1043,U+1044,U+1045,U+1046,U+1047,U+1048,U+1049'

execute('rm -rf ' + fonts_destination_dir)
execute('mkdir -p ' + fonts_destination_dir)

captcha_fonts_template = """package main

import (
\t_ "embed"
)

// this file auto generate at {generate_time}
// total files: {total_files}
// total size: {total_size}KB

{embeds}

{var_list_all}
"""

embed_template = """//go:embed {path}\nvar {embed_var} []byte"""
var_list_template = """// CaptchaFonts_{lang} list of embed fonts for {lang}\nvar CaptchaFonts_{lang} = [][]byte{{{embed_var}}}"""
lang_index = {}
embeds = []
var_list = {}
total_size = 0
total_files = 0
languages = os.listdir(fonts_source_dir)
for lang in languages:
  lang_dir = fonts_source_dir + "/" + lang

  for path in Path(lang_dir).rglob('*.ttf'):
    if lang not in lang_index:
      lang_index[lang] = -1

    if lang not in var_list:
      var_list[lang] = []

    total_files += 1
    lang_index[lang] += 1
    source_path = lang_dir + "/" + path.name
    md5_file = execute('md5sum ' + source_path + " | cut -d ' ' -f 1").strip()
    embed_var = lang + '_' + md5_file
    var_list[lang].append(embed_var)
    destination_name = embed_var + '.ttf'
    destination_path = fonts_destination_dir + '/' + destination_name

    source_file_size = os.path.getsize(source_path)
    print(source_path)
    command = 'pyftsubset {source_path} --name-IDs="" --unicodes={unicode_chars} --output-file={destination_path}'.format(
      source_path=source_path,
      unicode_chars=unicode_chars,
      destination_path=destination_path,
    )
    execute(command)

    embeds.append(embed_template.format(
      embed_var=embed_var,
      path=destination_path.replace(project_path + "/", "")),
    )
    embeds.append("")

    destination_file_size = os.path.getsize(destination_path)
    total_size += destination_file_size
    decrease_percent = 100 - round(destination_file_size / source_file_size * 100)
    print("Decrease size: "+ str(decrease_percent) + "%")


var_list_all = []
for lang, embed_var in var_list.items():
  var_list_all.append(var_list_template.format(lang=lang,embed_var=",".join(embed_var)))
  var_list_all.append("")

captcha_fonts = captcha_fonts_template.format(
  generate_time=datetime.datetime.now().strftime("%Y/%m/%d %H:%M"),
  total_files=total_files,
  total_size=round(total_size / 1024),
  var_list_all="\n".join(var_list_all),
  embeds="\n".join(embeds),
)

with open(project_path + "/captcha_fonts_list.go", 'w') as captcha_fonts_file:
  captcha_fonts_file.write(captcha_fonts)
